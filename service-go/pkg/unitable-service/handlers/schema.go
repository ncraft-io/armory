package handlers

import (
	"context"
	"errors"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// tableIdentity accepts either a table name or its database-qualified ID.
func tableIdentity(database, name string) (string, string, error) {
	if database == "" || strings.ContainsAny(database, ".%\"' ;()") {
		return "", "", core.NewInvalidArgumentError("invalid database name")
	}
	name = strings.TrimPrefix(name, database+".")
	if !nameRegex.MatchString(name) {
		return "", "", core.NewInvalidArgumentError("invalid table name %q", name)
	}
	return database + "." + name, name, nil
}

func (s unitableServer) getTable(ctx context.Context, database, name string) (*unitable.Table, error) {
	id, _, err := tableIdentity(database, name)
	if err != nil {
		return nil, err
	}
	table, err := model.GetTableModel().Get(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, core.NewNotFoundError("table %s not found", id)
	}
	return table, err
}

func validateColumns(table *unitable.Table) error {
	names, ids, fields := map[string]bool{}, map[string]bool{}, map[string]bool{}
	for _, col := range table.Columns {
		if col == nil || !nameRegex.MatchString(col.Name) || !col.IsTypeValid() {
			return core.NewInvalidArgumentError("invalid column name or type")
		}
		field := strcase.ToCamel(col.Name)
		if names[col.Name] || fields[field] || (col.Id != "" && ids[col.Id]) {
			return core.NewInvalidArgumentError("duplicate column name or id: %s", col.Name)
		}
		if col.Database != "" && col.Database != table.Database {
			return core.NewInvalidArgumentError("column database differs from table")
		}
		if col.TableId != "" && col.TableId != table.Id {
			return core.NewInvalidArgumentError("column belongs to another table")
		}
		if col.Repeated && (col.Type != "string" || col.Format != "") {
			return core.NewInvalidArgumentError("only plain string columns support repeated values")
		}
		names[col.Name], ids[col.Id], fields[field] = true, true, true
	}
	return nil
}

func prepareColumns(table *unitable.Table) {
	for _, col := range table.Columns {
		if col.Id == "" {
			col.Id = ksuid.New().String()
		}
		col.Database, col.TableId = table.Database, table.Id
		if col.CreateTime == nil {
			col.CreateTime = core.Now()
		}
		col.UpdateTime = core.Now()
	}
}

// Store the table and its has-many columns together. Columns omitted by a
// caller of this helper are removed, so schema deletion cannot leave ghosts.
func (s unitableServer) saveTable(ctx context.Context, table *unitable.Table) error {
	_ = model.GetColumnModel()
	err := model.GetTableModel().DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, col := range table.Columns {
			var count int64
			if err := tx.Model(&unitable.Column{}).Where("id = ? AND table_id <> ?", col.Id, table.Id).Count(&count).Error; err != nil {
				return err
			}
			if count != 0 {
				return core.NewInvalidArgumentError("column id %s belongs to another table", col.Id)
			}
		}
		if err := tx.Omit(clause.Associations).Clauses(clause.OnConflict{UpdateAll: true}).Create(table).Error; err != nil {
			return err
		}
		ids := make([]string, 0, len(table.Columns))
		for _, col := range table.Columns {
			ids = append(ids, col.Id)
		}
		stale := tx.Where("table_id = ?", table.Id)
		if len(ids) > 0 {
			stale = stale.Where("id NOT IN ?", ids)
		}
		if err := stale.Delete(&unitable.Column{}).Error; err != nil {
			return err
		}
		if len(table.Columns) > 0 {
			return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(table.Columns).Error
		}
		return nil
	})
	if err == nil {
		s.Synchro().InvalidateTable(table.Id)
	}
	return err
}

func (s unitableServer) CreateTable(ctx context.Context, in *pb.CreateTableRequest) (*unitable.Table, error) {
	if in == nil || in.Table == nil {
		return nil, core.NewInvalidArgumentError("not set the table")
	}
	s.instance.schemaMu.Lock()
	defer s.instance.schemaMu.Unlock()
	table := proto.Clone(in.Table).(*unitable.Table)
	if table.Database == "" {
		table.Database = in.Database
	}
	if in.Database != "" && in.Database != table.Database {
		return nil, core.NewInvalidArgumentError("database differs from request path")
	}
	id, name, err := tableIdentity(table.Database, table.Name)
	if err != nil {
		return nil, err
	}
	if table.Id != "" && table.Id != id {
		return nil, core.NewInvalidArgumentError("table id differs from database and name")
	}
	table.Id, table.Name = id, name
	if len(table.Columns) == 0 {
		return nil, core.NewInvalidArgumentError("table has no columns")
	}
	if err := validateColumns(table); err != nil {
		return nil, err
	}
	old, err := model.GetTableModel().Get(ctx, id)
	if err == nil {
		return s.updateTable(ctx, table, old, false)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	prepareColumns(table)
	table.CreateTime, table.UpdateTime = core.Now(), core.Now()
	if err := s.checkColumnIDs(ctx, table); err != nil {
		return nil, err
	}
	if err := s.Synchro().MigrateTable(ctx, table, nil, nil); err != nil {
		return nil, err
	}
	if err := s.saveTable(ctx, table); err != nil {
		return nil, err
	}
	return table, nil
}

func (s unitableServer) checkColumnIDs(ctx context.Context, table *unitable.Table) error {
	m := model.GetColumnModel()
	for _, col := range table.Columns {
		old, err := m.Get(ctx, col.Id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if old != nil && old.TableId != table.Id {
			return core.NewInvalidArgumentError("column id %s belongs to another table", col.Id)
		}
	}
	return nil
}

// Missing names/types identify an existing column by ID; booleans remain
// caller-controlled, so flags can be switched off.
func MergeColumn(target, src *unitable.Column) {
	if target.Name == "" {
		target.Name = src.Name
	}
	if target.Type == "" {
		target.Type, target.Format, target.Repeated = src.Type, src.Format, src.Repeated
	}
	if target.CreateTime == nil {
		target.CreateTime = src.CreateTime
	}
}

func (s unitableServer) updateTable(ctx context.Context, table, old *unitable.Table, force bool) (*unitable.Table, error) {
	renamed := map[string]string{}
	var drop []string
	byID, byName := map[string]*unitable.Column{}, old.ColumnIndex()
	for _, col := range old.Columns {
		byID[col.Id] = col
	}
	seen := map[string]bool{}
	for _, col := range table.Columns {
		if col == nil {
			return nil, core.NewInvalidArgumentError("nil column")
		}
		c := byID[col.Id]
		if n := byName[col.Name]; n != nil {
			if col.Id != "" && col.Id != n.Id {
				return nil, core.NewInvalidArgumentError("column name and id refer to different columns")
			}
			c = n
		}
		if c != nil {
			col.Id = c.Id
			MergeColumn(col, c)
			if col.Type != c.Type || col.Format != c.Format || col.Repeated != c.Repeated {
				if !force {
					return nil, core.NewInvalidArgumentError("changing column %s type requires force", c.Name)
				}
				drop = append(drop, c.Name)
			} else if col.Name != c.Name {
				renamed[c.Name] = col.Name
			}
			seen[c.Id] = true
		}
	}
	// Table updates are additive. Column deletion has its own explicit API.
	for _, col := range old.Columns {
		if !seen[col.Id] {
			table.Columns = append(table.Columns, proto.Clone(col).(*unitable.Column))
		}
	}
	if err := validateColumns(table); err != nil {
		return nil, err
	}
	prepareColumns(table)
	table.CreateTime, table.UpdateTime = old.CreateTime, core.Now()
	if err := s.checkColumnIDs(ctx, table); err != nil {
		return nil, err
	}
	if err := s.Synchro().MigrateTable(ctx, table, renamed, drop); err != nil {
		return nil, err
	}
	if err := s.saveTable(ctx, table); err != nil {
		return nil, err
	}
	return table, nil
}

func (s unitableServer) UpdateTable(ctx context.Context, in *pb.UpdateTableRequest) (*core.Null, error) {
	if in == nil || in.Table == nil {
		return nil, core.NewInvalidArgumentError("not set the table")
	}
	s.instance.schemaMu.Lock()
	defer s.instance.schemaMu.Unlock()
	table := proto.Clone(in.Table).(*unitable.Table)
	database := in.Database
	if database == "" {
		database = table.Database
	}
	if table.Database != "" && table.Database != database {
		return nil, core.NewInvalidArgumentError("database differs from request path")
	}
	name := in.Id
	if name == "" {
		name = table.Id
	}
	if name == "" {
		name = table.Name
	}
	old, err := s.getTable(ctx, database, name)
	if err != nil {
		return nil, err
	}
	if table.Id != "" && table.Id != old.Id && table.Id != old.Name {
		return nil, core.NewInvalidArgumentError("table id differs from request path")
	}
	if table.Name != "" && table.Name != old.Name {
		return nil, core.NewInvalidArgumentError("table rename is not supported")
	}
	table.Id, table.Name, table.Database = old.Id, old.Name, old.Database
	if len(table.Columns) == 0 {
		table.Columns = old.Columns
	}
	if _, err := s.updateTable(ctx, table, old, in.Force); err != nil {
		return nil, err
	}
	return &core.Null{}, nil
}

func (s unitableServer) GetTable(ctx context.Context, in *pb.GetTableRequest) (*unitable.Table, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	return s.getTable(ctx, in.Database, in.Id)
}

func (s unitableServer) ListTables(ctx context.Context, in *pb.ListTablesRequest) (*pb.ListTablesResponse, error) {
	if in == nil || in.Database == "" {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	qry, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	qry.AddFieldQuery("database", in.Database)
	tables, err := model.GetTableModel().List(ctx, qry)
	if err != nil {
		return nil, err
	}
	allQuery, _ := ParseQuery(in)
	allQuery.PageSize, allQuery.PageToken, allQuery.Skip = 0, "", 0
	allQuery.AddFieldQuery("database", in.Database)
	all, err := model.GetTableModel().List(ctx, allQuery)
	if err != nil {
		return nil, err
	}
	return &pb.ListTablesResponse{Tables: tables, TotalCount: int32(len(all)), NextPageToken: nextPageToken(qry, len(all))}, nil
}

func (s unitableServer) DeleteTable(ctx context.Context, in *pb.DeleteTableRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	s.instance.schemaMu.Lock()
	defer s.instance.schemaMu.Unlock()
	table, err := s.getTable(ctx, in.Database, in.Id)
	if err != nil {
		return nil, err
	}
	if in.Force {
		if err := s.Synchro().DropTable(ctx, table); err != nil {
			return nil, err
		}
	}
	err = model.GetTableModel().DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("table_id = ?", table.Id).Delete(&unitable.Column{}).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", table.Id).Delete(&unitable.Table{}).Error
	})
	if err != nil {
		return nil, err
	}
	s.Synchro().InvalidateTable(table.Id)
	return &core.Null{}, nil
}

func (s unitableServer) SyncTable(ctx context.Context, in *pb.SyncTableRequest) (*unitable.Table, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	s.instance.schemaMu.Lock()
	defer s.instance.schemaMu.Unlock()
	id, name, err := tableIdentity(in.Database, in.Id)
	if err != nil {
		return nil, err
	}
	d := synchro.GetDataDB(in.Database)
	if d == nil {
		return nil, core.NewNotFoundError("database not found")
	}
	if !d.WithContext(ctx).Migrator().HasTable(name) {
		return nil, core.NewNotFoundError("table %s not found", name)
	}
	types, err := d.WithContext(ctx).Migrator().ColumnTypes(name)
	if err != nil {
		return nil, err
	}
	table, err := model.GetTableModel().Get(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if table == nil {
		table = &unitable.Table{Id: id, Name: name, Database: in.Database, CreateTime: core.Now()}
	}
	old := table.ColumnIndex()
	table.Columns = nil
	for _, typ := range types {
		col := &unitable.Column{Name: typ.Name()}
		if previous := old[col.Name]; previous != nil {
			col = proto.Clone(previous).(*unitable.Column)
		}
		col.Type, col.Format, col.Repeated, err = columnType(typ.DatabaseTypeName())
		if err != nil {
			return nil, err
		}
		table.Columns = append(table.Columns, col)
	}
	if err := validateColumns(table); err != nil {
		return nil, err
	}
	prepareColumns(table)
	table.UpdateTime = core.Now()
	if err := s.saveTable(ctx, table); err != nil {
		return nil, err
	}
	return table, nil
}

func columnType(name string) (string, string, bool, error) {
	n := strings.ToLower(name)
	switch {
	case n == "_text" || n == "_varchar" || n == "text[]" || n == "varchar[]":
		return "string", "", true, nil
	case strings.Contains(n, "bool"):
		return "bool", "", false, nil
	case strings.Contains(n, "int") || strings.Contains(n, "serial"):
		return "integer", "", false, nil
	case strings.Contains(n, "float") || strings.Contains(n, "double") || strings.Contains(n, "real") || strings.Contains(n, "numeric") || strings.Contains(n, "decimal"):
		return "float", "", false, nil
	case strings.Contains(n, "time") || n == "date":
		return "string", "datetime", false, nil
	case strings.Contains(n, "json"):
		return "string", "json", false, nil
	case strings.Contains(n, "geom"):
		return "string", "geometry", false, nil
	case strings.Contains(n, "char") || strings.Contains(n, "text") || n == "uuid":
		return "string", "", false, nil
	case strings.Contains(n, "blob") || strings.Contains(n, "binary") || n == "bytea":
		return "string", "bytes", false, nil
	default:
		return "", "", false, core.NewInvalidArgumentError("unsupported database column type %q", name)
	}
}

func (s unitableServer) ListDatabases(ctx context.Context, in *pb.ListDatabasesRequest) (*pb.ListDatabasesResponse, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	qry, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	if err := qry.Normalize(); err != nil {
		return nil, core.NewInvalidArgumentError("invalid query: %s", err)
	}
	tables, err := model.GetTableModel().List(ctx, nil)
	if err != nil {
		return nil, err
	}
	groups := map[string]*unitable.Database{}
	cfg := &synchro.DataDBCfg{}
	if err := config.ScanFrom(cfg, "dataDB"); err != nil {
		return nil, err
	}
	for name := range cfg.Dbs {
		groups[name] = &unitable.Database{Name: name}
	}
	for _, table := range tables {
		if groups[table.Database] == nil {
			groups[table.Database] = &unitable.Database{Name: table.Database}
		}
		groups[table.Database].Tables = append(groups[table.Database].Tables, table)
	}
	if len(groups) == 0 {
		return &pb.ListDatabasesResponse{}, nil
	}
	names := make([]string, 0, len(groups))
	for name := range groups {
		names = append(names, name)
	}
	sort.Strings(names)
	// Build a bound catalog relation so the standard query language applies to
	// database names, not to the names of tables inside them.
	var selects []string
	var args []interface{}
	for _, name := range names {
		selects = append(selects, "SELECT ? AS name, ? AS display_name, ? AS description")
		args = append(args, name, groups[name].DisplayName, groups[name].Description)
	}
	d := model.GetTableModel().DB.WithContext(ctx)
	relation := d.Raw(strings.Join(selects, " UNION ALL "), args...)
	fields := query.Fields{"name": {Type: query.FieldTypeString}, "display_name": {Type: query.FieldTypeString}, "description": {Type: query.FieldTypeString}}
	if qry.Filter != nil {
		if _, _, err := query.GenerateExpressionQuery(d.Dialector.Name(), qry.Filter, fields); err != nil {
			return nil, core.NewInvalidArgumentError("invalid database filter: %s", err)
		}
	}
	base := func() *gorm.DB { return d.Table("(?) AS unitable_databases", relation) }
	allQuery, _ := ParseQuery(in)
	allQuery.PageSize, allQuery.PageToken, allQuery.Skip = 0, "", 0
	if err := allQuery.Normalize(); err != nil {
		return nil, err
	}
	var count int64
	if err := d.Table("(?) AS matched_databases", allQuery.Apply(base(), fields)).Count(&count).Error; err != nil {
		return nil, err
	}
	type databaseRow struct {
		Name        string
		DisplayName string
		Description string
	}
	var rows []databaseRow
	if err := qry.Apply(base(), fields).Order("name").Scan(&rows).Error; err != nil {
		return nil, err
	}
	databases := make([]*unitable.Database, 0, len(rows))
	for _, row := range rows {
		database := &unitable.Database{Name: row.Name, DisplayName: row.DisplayName, Description: row.Description}
		if group := groups[row.Name]; group != nil {
			database.Tables = group.Tables
		}
		databases = append(databases, database)
	}
	return &pb.ListDatabasesResponse{Databases: databases, TotalCount: int32(count), NextPageToken: nextPageToken(qry, int(count))}, nil
}

// Used by column listings to keep database/table scope independent of filters.
func scopedColumnQuery(in *pb.ListColumnsRequest, id string) (*query.Query, error) {
	qry, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	qry.AddFieldQuery("database", in.Database).AddFieldQuery("table_id", id)
	return qry, nil
}
