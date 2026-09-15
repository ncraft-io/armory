package handlers

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var queryNameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,127}$`)

func queryIdentity(database, name string) (string, string, error) {
	if !queryNameRegex.MatchString(database) {
		return "", "", core.NewInvalidArgumentError("invalid database name")
	}
	if synchro.GetDataDB(database) == nil {
		return "", "", core.NewNotFoundError("database %s not found", database)
	}
	name = strings.TrimPrefix(name, database+".")
	if !queryNameRegex.MatchString(name) {
		return "", "", core.NewInvalidArgumentError("invalid query name")
	}
	return database + "." + name, name, nil
}
func queryConfigKey(q *unitable.DbQuery) string {
	if q.Database != "" {
		return q.Database + "." + q.Name
	}
	return q.Name
}
func storedQuery(ctx context.Context, database, name string) (*unitable.DbQuery, error) {
	_, name, err := queryIdentity(database, name)
	if err != nil {
		return nil, err
	}
	q := &unitable.DbQuery{}
	tx := model.GetDbQueryModel().DB.WithContext(ctx).Where(clause.Eq{Column: "database", Value: database}).Where("name = ?", name).First(q)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, core.NewNotFoundError("query %s not found", name)
	}
	return q, tx.Error
}
func (s unitableServer) resolveQuery(ctx context.Context, database, name string) (*unitable.DbQuery, error) {
	id, name, err := queryIdentity(database, name)
	if err != nil {
		return nil, err
	}
	q := s.Queries()[id]
	if q == nil {
		q = s.Queries()[name]
	}
	if q != nil && (q.Database == "" || q.Database == database) {
		return proto.Clone(q).(*unitable.DbQuery), nil
	}
	return storedQuery(ctx, database, name)
}

func validateQueryColumns(q *unitable.DbQuery) error {
	if len(q.Columns) == 0 {
		return core.NewInvalidArgumentError("query result columns are required")
	}
	fields := map[string]bool{}
	for _, col := range q.Columns {
		if col == nil || !nameRegex.MatchString(col.Name) || !col.IsTypeValid() {
			return core.NewInvalidArgumentError("invalid result column name or type")
		}
		field := strcase.ToCamel(col.Name)
		if fields[field] {
			return core.NewInvalidArgumentError("duplicate result column %s", col.Name)
		}
		fields[field] = true
		if col.Repeated && (col.Type != "string" || col.Format != "") {
			return core.NewInvalidArgumentError("only plain string result columns support repeated values")
		}
	}
	if q.JsonStyle == "" {
		q.JsonStyle = synchro.LowerCamel
	}
	switch strings.ToLower(q.JsonStyle) {
	case "snake":
		q.JsonStyle = "snake"
	case "lowercamel", "lower_camel":
		q.JsonStyle = synchro.LowerCamel
	default:
		return core.NewInvalidArgumentError("unsupported json_style")
	}
	return nil
}

// validateDbQuery checks plans without ANALYZE and inspects result metadata
// through a zero-row SELECT. It never fetches business query results.
func validateDbQuery(ctx context.Context, database string, input *unitable.DbQuery) (*unitable.DbQuery, error) {
	if input == nil {
		return nil, core.NewInvalidArgumentError("query is required")
	}
	q := proto.Clone(input).(*unitable.DbQuery)
	id, name, err := queryIdentity(database, q.Name)
	if err != nil {
		return nil, err
	}
	if q.Name != name || (q.Id != "" && q.Id != id) || (q.Database != "" && q.Database != database) {
		return nil, core.NewInvalidArgumentError("query identity differs from request path")
	}
	q.Id, q.Database = id, database
	if err = validateQueryColumns(q); err != nil {
		return nil, err
	}
	template, err := parseQueryTemplate(q.Sql)
	if err != nil {
		return nil, core.NewInvalidArgumentError("invalid SQL: %s", err)
	}
	sample := map[string]interface{}{}
	definitions := map[string]*unitable.DbQuery_Parameter{}
	for _, p := range q.Parameters {
		if p == nil || !queryNameRegex.MatchString(p.Name) || parameterType(p) == "" || p.PgArray && !p.IsArray {
			return nil, core.NewInvalidArgumentError("invalid query parameter definition")
		}
		if previous := definitions[p.Name]; previous != nil && (parameterType(previous) != parameterType(p) || previous.IsArray != p.IsArray || previous.PgArray != p.PgArray) {
			return nil, core.NewInvalidArgumentError("inconsistent repeated parameter %s", p.Name)
		}
		definitions[p.Name] = p
		var v interface{} = "validation"
		switch parameterType(p) {
		case "integer":
			v = int64(1)
		case "float":
			v = float64(1)
		case "bool":
			v = true
		}
		if p.IsArray {
			v = []interface{}{v}
		}
		sample[p.Name] = v
	}
	if template.count != len(q.Parameters) {
		return nil, core.NewInvalidArgumentError("SQL placeholder count differs from parameter definitions")
	}
	d := synchro.GetDataDB(database)
	dialect := d.Dialector.Name()
	switch dialect {
	case "sqlite", "mysql", "postgres":
	default:
		return nil, core.NewInvalidArgumentError("query validation is unsupported for database dialect %s", dialect)
	}
	sqlDB, err := d.DB.DB()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// Validate every reachable combination of optional parameter presence. Repeated
	// names share a switch, so legacy repeated optional conditions stay inexpensive.
	required := map[string]bool{}
	optional := []string{}
	seen := map[string]bool{}
	n := 0
	for _, part := range template.parts {
		for _, token := range part.tokens {
			if token.parameter {
				name := q.Parameters[n].Name
				n++
				if !part.optional {
					required[name] = true
				} else if !seen[name] {
					optional = append(optional, name)
					seen[name] = true
				}
			}
		}
	}
	switches := optional[:0]
	for _, name := range optional {
		if !required[name] {
			switches = append(switches, name)
		}
	}
	if len(switches) > 6 {
		return nil, core.NewInvalidArgumentError("at most 6 independently optional parameters are supported")
	}
	for mask := 0; mask < 1<<len(switches); mask++ {
		values := map[string]interface{}{}
		for k, v := range sample {
			values[k] = v
		}
		for bit, name := range switches {
			if mask&(1<<bit) == 0 {
				delete(values, name)
			}
		}
		sql, args, bindErr := template.bind(q, values, dialect)
		if bindErr != nil {
			if mask == (1<<len(switches))-1 {
				return nil, core.NewInvalidArgumentError("invalid parameters: %s", bindErr)
			}
			continue
		} // partial optional groups are rejected on execution too
		// NULL bindings let the database infer casts (UUID, dates, etc.) without
		// rejecting valid SQL because an invented sample value cannot be cast.
		args = make([]interface{}, len(args))
		rows, explainErr := sqlDB.QueryContext(ctx, "EXPLAIN "+sql, args...)
		if explainErr != nil {
			return nil, core.NewInvalidArgumentError("SQL validation failed: %s", explainErr)
		}
		for rows.Next() {
		}
		explainErr = rows.Err()
		closeErr := rows.Close()
		if explainErr != nil {
			return nil, core.NewInvalidArgumentError("SQL validation failed: %s", explainErr)
		}
		if closeErr != nil {
			return nil, closeErr
		}
		metadata, metadataErr := sqlDB.QueryContext(ctx, "SELECT * FROM ("+sql+"\n) AS db_query_validation WHERE 1=0", args...)
		if metadataErr != nil {
			return nil, core.NewInvalidArgumentError("invalid result query: %s", metadataErr)
		}
		columns, columnsErr := metadata.Columns()
		closeErr = metadata.Close()
		if columnsErr != nil {
			return nil, columnsErr
		}
		if closeErr != nil {
			return nil, closeErr
		}
		if len(columns) != len(q.Columns) {
			return nil, core.NewInvalidArgumentError("result column count differs from query definition")
		}
		for i, name := range columns {
			if name != q.Columns[i].Name {
				return nil, core.NewInvalidArgumentError("result column %d is %s, expected %s", i+1, name, q.Columns[i].Name)
			}
		}
	}
	return q, nil
}

func (s unitableServer) CreateDbQuery(ctx context.Context, in *pb.CreateDbQueryRequest) (*unitable.DbQuery, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	q, err := validateDbQuery(ctx, in.Database, in.Query)
	if err != nil {
		return nil, err
	}
	var existing int64
	if err := model.GetDbQueryModel().DB.WithContext(ctx).Model(&unitable.DbQuery{}).Where(clause.Eq{Column: "database", Value: q.Database}).Where("name = ?", q.Name).Count(&existing).Error; err != nil {
		return nil, err
	}
	if existing != 0 {
		return nil, core.NewAlreadyExistsError("query %s already exists", q.Name)
	}
	q.CreateTime, q.UpdateTime = core.Now(), core.Now()
	// Plain INSERT never overwrites a definition on a duplicate create request.
	tx := model.GetDbQueryModel().DB.WithContext(ctx).Create(q)
	if tx.Error != nil {
		// Some MySQL configurations report a no-op UPSERT as a successful
		// insert. Use INSERT and inspect a conflicting ID to keep duplicate
		// creates consistent across drivers, including concurrent requests.
		var conflicts int64
		lookupErr := model.GetDbQueryModel().DB.WithContext(ctx).Model(&unitable.DbQuery{}).Where("id = ?", q.Id).Count(&conflicts).Error
		if lookupErr == nil && conflicts > 0 {
			return nil, core.NewAlreadyExistsError("query %s already exists", q.Name)
		}
		return nil, tx.Error
	}
	return q, nil
}
func (s unitableServer) UpdateDbQuery(ctx context.Context, in *pb.UpdateDbQueryRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	old, err := storedQuery(ctx, in.Database, in.Id)
	if err != nil {
		return nil, err
	}
	if in.Query == nil {
		return nil, core.NewInvalidArgumentError("query is required")
	}
	input := proto.Clone(in.Query).(*unitable.DbQuery)
	if input.Id != "" && input.Id != old.Id {
		return nil, core.NewInvalidArgumentError("query id cannot be changed")
	}
	input.Id = ""
	q, err := validateDbQuery(ctx, in.Database, input)
	if err != nil {
		return nil, err
	}
	if q.Name != old.Name {
		return nil, core.NewInvalidArgumentError("query name cannot be changed")
	}
	q.Id, q.CreateTime, q.UpdateTime = old.Id, old.CreateTime, core.Now()
	count, err := model.GetDbQueryModel().UpdateAll(ctx, q)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, core.NewNotFoundError("query no longer exists")
	}
	return &core.Null{}, nil
}
func (s unitableServer) GetDbQuery(ctx context.Context, in *pb.GetDbQueryRequest) (*unitable.DbQuery, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	return storedQuery(ctx, in.Database, in.Id)
}
func (s unitableServer) ListDbQueries(ctx context.Context, in *pb.ListDbQueriesRequest) (*pb.ListDbQueriesResponse, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if _, _, err := queryIdentity(in.Database, "query"); err != nil {
		return nil, err
	}
	qry, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	qry.AddFieldQuery("database", in.Database)
	values, err := model.GetDbQueryModel().List(ctx, qry)
	if err != nil {
		return nil, err
	}
	allQuery, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	allQuery.PageSize, allQuery.PageToken, allQuery.Skip = 0, "", 0
	allQuery.AddFieldQuery("database", in.Database)
	all, err := model.GetDbQueryModel().List(ctx, allQuery)
	if err != nil {
		return nil, err
	}
	return &pb.ListDbQueriesResponse{DbQueries: values, TotalCount: int32(len(all)), NextPageToken: nextPageToken(qry, len(all))}, nil
}
func (s unitableServer) DeleteDbQuery(ctx context.Context, in *pb.DeleteDbQueryRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	q, err := storedQuery(ctx, in.Database, in.Id)
	if err != nil {
		return nil, err
	}
	count, err := model.GetDbQueryModel().Delete(ctx, q.Id)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, core.NewNotFoundError("query no longer exists")
	}
	return &core.Null{}, nil
}
func (s unitableServer) runQuery(ctx context.Context, database, name string, values map[string]interface{}) ([]*core.Object, error) {
	q, err := s.resolveQuery(ctx, database, name)
	if err != nil {
		return nil, err
	}
	if err = validateQueryColumns(q); err != nil {
		return nil, err
	}
	if strings.TrimSpace(q.Sql) == "" {
		return []*core.Object{q.Example()}, nil
	} // legacy configuration examples
	t, err := parseQueryTemplate(q.Sql)
	if err != nil {
		return nil, core.NewInvalidArgumentError("invalid query SQL: %s", err)
	}
	sql, args, err := t.bind(q, values, synchro.GetDataDB(database).Dialector.Name())
	if err != nil {
		return nil, core.NewInvalidArgumentError("invalid query parameters: %s", err)
	}
	q.Sql = sql
	return s.Synchro().QueryPrepared(ctx, database, q.Name, q, args)
}
func (s unitableServer) RunDbQuery(ctx context.Context, in *pb.RunDbQueryRequest) (*pb.RunDbQueryResponse, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	var values map[string]interface{}
	if in.Parameters != nil {
		values = in.Parameters.ToMapInterface()
	}
	qry, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	if in.Filter != "" || in.Order != nil || in.FieldMask != nil || in.Unique {
		return nil, core.NewInvalidArgumentError("predefined SQL controls filtering, ordering and result columns")
	}
	objects, err := s.runQuery(ctx, in.Database, in.Id, values)
	if err != nil {
		return nil, err
	}
	total := len(objects)
	start := int64(qry.Skip)
	page, _ := strconv.ParseInt(qry.PageToken, 10, 32)
	if qry.PageSize > 0 {
		start += page * int64(qry.PageSize)
	}
	if start > int64(total) {
		start = int64(total)
	}
	end := int64(total)
	if qry.PageSize > 0 && start+int64(qry.PageSize) < end {
		end = start + int64(qry.PageSize)
	}
	return &pb.RunDbQueryResponse{Objects: objects[start:end], TotalCount: int32(total), NextPageToken: nextPageToken(qry, total)}, nil
}
