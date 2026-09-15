package synchro

import (
	"context"
	"database/sql"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"sync"
)

type MetaTable struct {
	Table      *unitable.Table
	Struct     *DynamicStruct
	FieldsInfo query.Fields
}

type Synchro struct {
	cacheMu sync.Mutex
	Tables  sync.Map //map[string]*MetaTable
}

func TableId(database string, table string) string {
	return strings.Join([]string{database, table}, ".")
}

func New() *Synchro {
	return &Synchro{}
}

func (s *Synchro) GetMetaTable(tableId string, table *unitable.Table) *MetaTable {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	if table == nil { // if table not nil, will using user table, and not cache
		if value, ok := s.Tables.Load(tableId); ok {
			if mt, ok := value.(*MetaTable); ok {
				return mt
			}
		}
	}

	tb := table
	if tb == nil {
		t, err := model.GetTableModel().Get(context.Background(), tableId)
		if err != nil {
			logs.ErrLogw("failed to get the table", "name", table, "error", err)
			return nil
		}
		if len(t.GetColumns()) == 0 {
			logs.ErrLogw("failed to get any columns in the table", "name", table)
			return nil
		}

		tb = t
	}

	if GetDataDB(tb.Database) == nil {
		return nil
	}
	meta := &MetaTable{
		Table:      tb,
		Struct:     NewDynamicStruct(tb),
		FieldsInfo: make(query.Fields),
	}
	for _, col := range tb.Columns {
		meta.FieldsInfo[col.Name] = col.ToFieldInfo()
	}

	logs.Infow("get the meta table ok", "database", tb.Database, "table", tb.Name, "columns", len(tb.Columns))

	if table == nil {
		s.Tables.Store(tableId, meta)
	}
	return meta
}

func (s *Synchro) InvalidateTable(id string) {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	s.Tables.Delete(id)
}

func (s *Synchro) CreateTable(ctx context.Context, table *unitable.Table) error {
	return s.MigrateTable(ctx, table, nil, nil)
}

func (s *Synchro) MigrateTable(ctx context.Context, table *unitable.Table, renamedCols map[string]string, dropCols []string) error {
	d := GetDataDB(table.Database)
	if d == nil {
		return core.NewNotFoundError("database %s not found", table.Database)
	}
	// Some dialects implicitly commit DDL. Validate the complete request before
	// entering here; always invalidate cached structs even after partial DDL.
	defer s.InvalidateTable(table.Id)
	tx := d.WithContext(ctx).Table(table.Name)
	// AutoMigrate adds indexes but does not remove them when a flag is cleared.
	if tx.Migrator().HasTable(table.Name) {
		if old, err := model.GetTableModel().Get(ctx, table.Id); err == nil {
			oldMeta := s.GetMetaTable(table.Id, old)
			statement := &gorm.Statement{DB: tx}
			if err := statement.ParseWithSpecialTableName(oldMeta.Struct.New(), table.Name); err != nil {
				return err
			}
			current := map[string]*unitable.Column{}
			for _, col := range table.Columns {
				current[col.Id] = col
			}
			byName := old.ColumnIndex()
			for _, index := range statement.Schema.ParseIndexes() {
				if len(index.Fields) != 1 {
					continue
				}
				previous := byName[index.Fields[0].DBName]
				if previous == nil {
					continue
				}
				next := current[previous.Id]
				remove := next == nil || next.Name != previous.Name
				if next != nil {
					if index.Class == "UNIQUE" {
						remove = remove || !next.Unique
					} else {
						remove = remove || !next.Indexed
					}
				}
				if remove && tx.Migrator().HasIndex(oldMeta.Struct.New(), index.Name) {
					if err := tx.Migrator().DropIndex(oldMeta.Struct.New(), index.Name); err != nil {
						return err
					}
				}
			}
		}
	}
	if len(renamedCols) > 0 || len(dropCols) > 0 {
		meta := s.GetMetaTable(table.Id, nil)
		if meta == nil {
			return core.NewNotFoundError("original table %s not found", table.Id)
		}
		obj := meta.Struct.New()
		for from, to := range renamedCols {
			if err := tx.Migrator().RenameColumn(obj, from, to); err != nil {
				return err
			}
		}
		for _, name := range dropCols {
			if err := tx.Migrator().DropColumn(obj, name); err != nil {
				return err
			}
		}
	}
	meta := s.GetMetaTable(table.Id, table)
	if err := tx.AutoMigrate(meta.Struct.New()); err != nil {
		return err
	}
	if d.Config.Driver == db.PostgresDriverName {
		if geom := table.GetGeometryColumn(); geom != nil {
			stmt := &gorm.Statement{DB: tx}
			sql := "CREATE INDEX IF NOT EXISTS " + stmt.Quote("idx_"+table.Name+"_geometry") + " ON " + stmt.Quote(table.Name) + " USING GIST (" + stmt.Quote(geom.Name) + ")"
			if err := tx.Exec(sql).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Synchro) DropTable(ctx context.Context, table *unitable.Table) error {
	d := GetDataDB(table.Database)
	if d == nil {
		return core.NewNotFoundError("database %s not found", table.Database)
	}
	defer s.InvalidateTable(table.Id)
	return d.WithContext(ctx).Migrator().DropTable(table.Name)
}

func (s *Synchro) GetRow(ctx context.Context, table string, id string) (*core.Object, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	row := meta.Struct.New()
	resul := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).First(row, "id = ?", id)
	if resul.Error != nil {
		return nil, core.NewNotFoundError("%s is not exist in %s, %s", id, table, resul.Error.Error())
	}

	obj, err := ParseObject(row)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *Synchro) BatchGetRow(ctx context.Context, table string, ids ...string) ([]*core.Object, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return nil, core.NewNotFoundError("the table %s is not exist", table)
	}
	if len(ids) == 0 {
		return nil, nil
	}

	rows := meta.Struct.NewSliceOf()
	resul := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Where("id IN ?", ids).Find(rows)
	if resul.Error != nil {
		return nil, core.NewNotFoundError("%s is not exist in %s, %s", ids, table, resul.Error.Error())
	}

	objs, err := ParseObjects(rows)
	if err != nil {
		return nil, err
	}

	return objs, nil
}

func (s *Synchro) QueryBy(ctx context.Context, database string, tableName string, query *unitable.DbQuery, arguments []interface{}) ([]*core.Object, error) {
	return s.queryBy(ctx, database, tableName, query, arguments, false)
}

// QueryPrepared executes SQL with native dialect placeholders, preserving literal
// question marks. Arguments must be bound separately by the caller.
func (s *Synchro) QueryPrepared(ctx context.Context, database, tableName string, query *unitable.DbQuery, arguments []interface{}) ([]*core.Object, error) {
	return s.queryBy(ctx, database, tableName, query, arguments, true)
}

func (s *Synchro) queryBy(ctx context.Context, database, tableName string, query *unitable.DbQuery, arguments []interface{}, prepared bool) ([]*core.Object, error) {

	table := &unitable.Table{
		Database:  database,
		Name:      tableName,
		JsonStyle: query.JsonStyle,
		Columns:   query.Columns,
	}
	meta := &MetaTable{
		Table:  table,
		Struct: NewDynamicStruct(table),
	}

	d := GetDataDB(database)
	if d == nil {
		return nil, core.NewNotFoundError("database %s not found", database)
	}
	var rows *sql.Rows
	var err error
	if prepared {
		native, nativeErr := d.DB.DB()
		if nativeErr != nil {
			return nil, nativeErr
		}
		rows, err = native.QueryContext(ctx, query.Sql, arguments...)
	} else {
		rows, err = d.WithContext(ctx).Table(table.Name).Raw(query.Sql, arguments...).Rows()
	}
	if err != nil {
		return nil, core.NewNotFoundError("failed to query the rows in %s, %s", table, err.Error())
	}

	defer rows.Close()
	var objs []*core.Object
	for rows.Next() {
		row := meta.Struct.New()
		if err = GetDataDB(database).ScanRows(rows, row); err != nil {
			return nil, core.NewInternalError("failed to scan the row in %s, %s", table, err.Error())
		}

		obj, err := ParseObject(row)
		if err != nil {
			return nil, core.NewInternalError("failed to parse the row in %s, %s", table, err.Error())
		}
		objs = append(objs, obj)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return objs, nil
}

func (s *Synchro) QueryRows(ctx context.Context, table string, query *query.Query) ([]*core.Object, int, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil || meta.Table == nil {
		logs.Warnw("the table is not exist", "table", table)
		return nil, 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name)
	if query != nil {
		tx = query.Apply(tx, meta.FieldsInfo)
	} else {
		tx = tx.Select("*")
	}

	rows, err := tx.Rows()
	if err != nil {
		return nil, 0, core.NewNotFoundError("failed to query the rows in %s, %s", table, err.Error())
	}

	st := NewDynamicStructWith(query, meta)

	defer rows.Close()
	var objs []*core.Object
	for rows.Next() {
		row := st.New()
		if err = GetDataDB(meta.Table.Database).ScanRows(rows, row); err != nil {
			return nil, 0, err
		}

		obj, err := ParseObject(row)
		if err != nil {
			return nil, 0, err
		}
		objs = append(objs, obj)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	if err := rows.Close(); err != nil {
		return nil, 0, err
	}
	totalCnt := len(objs)
	if query != nil && query.PageSize > 0 { // pagination
		// get the total count
		tx = GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name)
		if query != nil {
			countQuery := *query
			countQuery.Skip = 0
			tx = countQuery.TotalCount(tx, meta.FieldsInfo)
		} else {
			tx = tx.Select("COUNT(*)")
		}
		row := tx.Row()

		if row != nil {
			err = row.Scan(&totalCnt)
			if err != nil {
				return nil, 0, core.NewNotFoundError("failed to query the total count in %s, %s", table, err.Error())
			}
		}
	}

	return objs, totalCnt, nil
}

func (s *Synchro) InsertRow(ctx context.Context, table string, row *core.Object) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}
	logs.Debugf("the meta table is %v", meta.Table)

	data, err := meta.Struct.NewOf(row)
	if err != nil {
		return 0, err
	}
	logs.Debugf("the create row is %v", data)

	result := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Clauses(clause.OnConflict{UpdateAll: true}).Create(data)
	return result.RowsAffected, result.Error
}

func (s *Synchro) InsertRows(ctx context.Context, table string, rows ...*core.Object) (int64, []int, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	irs := make([]int, len(rows))

	// Continuous session mode
	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})
	err := tx.Transaction(func(tx *gorm.DB) error {
		for i, row := range rows {
			data, err := meta.Struct.NewOf(row)
			if err != nil {
				return err
			}
			result := tx.Table(meta.Table.Name).Clauses(clause.OnConflict{UpdateAll: true}).Create(data)
			if result.Error != nil {
				// return any error will roll back
				logs.Warnw("failed to create the row, will rollback", "row", row, "error", result.Error)
				return result.Error
			}
			if result.RowsAffected > 0 {
				irs[i] = 1
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return 0, nil, err
	}
	return int64(len(rows)), irs, nil
}

func (s *Synchro) UpdateRow(ctx context.Context, table string, row *core.Object) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	data, err := meta.Struct.NewOf(row)
	if err != nil {
		return 0, err
	}

	updateRow := FilterOutId(row.ToMapInterface(), meta.Table)
	result := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Model(data).Updates(updateRow)
	return result.RowsAffected, result.Error
}

func (s *Synchro) UpdateRows(ctx context.Context, table string, rows ...*core.Object) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	var datas []interface{}
	for _, row := range rows {
		r, err := meta.Struct.NewOf(row)
		if err != nil {
			return 0, err
		}
		datas = append(datas, r)
	}

	// Continuous session mode
	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})
	err := tx.Transaction(func(tx *gorm.DB) error {
		for i, data := range datas {
			updateRow := FilterOutId(rows[i].ToMapInterface(), meta.Table)
			if err := tx.Table(meta.Table.Name).Model(data).Updates(updateRow).Error; err != nil {
				// return any error will roll back
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	return int64(len(datas)), err
}

// UpdateInsertRows retains its historical name but updates only existing rows.
// It writes the supplied fields, including zero values, in one transaction.
func (s *Synchro) UpdateInsertRows(ctx context.Context, table string, rows ...*core.Object) (int64, []int, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, nil, core.NewNotFoundError("table %s not found", table)
	}
	irs := make([]int, len(rows))
	var affected int64
	err := GetDataDB(meta.Table.Database).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i, row := range rows {
			data, err := meta.Struct.NewOf(row)
			if err != nil {
				return err
			}
			values := FilterOutId(row.ToMapInterface(), meta.Table)
			result := tx.Table(meta.Table.Name).Model(data).Updates(values)
			if result.Error != nil {
				return result.Error
			}
			irs[i] = int(result.RowsAffected)
			affected += result.RowsAffected
		}
		return nil
	})
	if err != nil {
		return 0, nil, err
	}
	return affected, irs, nil
}

func (s *Synchro) DeleteRows(ctx context.Context, table string, ids ...string) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	if len(ids) == 0 {
		return 0, nil
	}
	result := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Where("id IN ?", ids).Delete(meta.Struct.New())
	return result.RowsAffected, result.Error
}

func (s *Synchro) CalcStats(ctx context.Context, table string, query *query.Query) ([]*core.Object, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name)
	if query == nil {
		return nil, core.NewInvalidArgumentError("the query is null from table %s", table)
	}

	tx = query.Apply(tx, meta.FieldsInfo)

	rows, err := tx.Rows()
	if err != nil {
		return nil, core.NewNotFoundError("failed to query the rows in %s, %s", table, err.Error())
	}

	st := NewDynamicStructWith(query, meta)

	defer rows.Close()
	var objs []*core.Object
	for rows.Next() {
		row := st.New()
		if err = GetDataDB(meta.Table.Database).ScanRows(rows, row); err != nil {
			return nil, err
		}

		obj, err := ParseObject(row)
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return objs, nil
}
