package handlers

import (
	"context"
	"sort"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// configuredQuery is shared by execution and effective reads. The returned
// definition belongs to the requested database; the startup configuration is
// never mutated, including its nested parameters and columns.
func (s unitableServer) configuredQuery(database, name string) *unitable.DbQuery {
	q := s.Queries()[database+"."+name]
	if q == nil {
		q = s.Queries()[name]
	}
	if q == nil || (q.Database != "" && q.Database != database) {
		return nil
	}
	result := proto.Clone(q).(*unitable.DbQuery)
	result.Id, result.Database = database+"."+name, database
	switch strings.ToLower(result.JsonStyle) {
	case "", "lowercamel", "lower_camel":
		result.JsonStyle = synchro.LowerCamel
	}
	return result
}

var effectiveQueryFields = query.Fields{
	"id": {Type: query.FieldTypeString}, "name": {Type: query.FieldTypeString},
	"sql": {Type: query.FieldTypeString}, "parameters": {Type: query.FieldTypeJSON, Repeated: true},
	"database": {Type: query.FieldTypeString}, "json_style": {Type: query.FieldTypeString},
	"columns":     {Type: query.FieldTypeJSON, Repeated: true},
	"create_time": {Type: query.FieldTypeDatetime}, "update_time": {Type: query.FieldTypeDatetime},
}

// effectiveQueryRelation forms a read-only derived table. Shadowed persisted
// rows are removed before unioning the winning configuration entries, so the
// existing query engine can filter, order and project the final definitions.
// Configuration values are SQL bindings, never SQL fragments or stored rows.
func (s unitableServer) effectiveQueryRelation(ctx context.Context, database string) *gorm.DB {
	names := map[string]bool{}
	for _, q := range s.Queries() {
		if q != nil && (q.Database == "" || q.Database == database) {
			names[q.Name] = true
		}
	}
	configured := make([]*unitable.DbQuery, 0, len(names))
	for name := range names {
		if q := s.configuredQuery(database, name); q != nil {
			configured = append(configured, q)
		}
	}
	sort.Slice(configured, func(i, j int) bool { return configured[i].Name < configured[j].Name })
	overridden := make([]string, 0, len(configured))
	for _, q := range configured {
		overridden = append(overridden, q.Name)
	}
	d := model.GetDbQueryModel().DB.WithContext(ctx)
	stored := d.Model(&unitable.DbQuery{}).Select([]string{"id", "name", "sql", "parameters", "database", "json_style", "columns", "create_time", "update_time"}).Where(clause.Eq{Column: "database", Value: database})
	if len(overridden) == 0 {
		return stored
	}
	stored = stored.Where("name NOT IN ?", overridden)
	sql := "?"
	args := []interface{}{stored}
	for _, q := range configured {
		sql += " UNION ALL SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?"
		args = append(args, q.Id, q.Name, q.Sql, q.Parameters, q.Database, q.JsonStyle, q.Columns, q.CreateTime, q.UpdateTime)
	}
	return d.Table("(?) AS effective_queries", d.Raw(sql, args...)).Model(&unitable.DbQuery{})
}

func (s unitableServer) listEffectiveDbQueries(ctx context.Context, in *pb.ListDbQueriesRequest, qry *query.Query) (*pb.ListDbQueriesResponse, error) {
	load := func(q *query.Query) ([]*unitable.DbQuery, error) {
		if err := q.Normalize(); err != nil {
			return nil, err
		}
		tx := q.Apply(s.effectiveQueryRelation(ctx, in.Database), effectiveQueryFields)
		// Name is unique in the merged view and makes page boundaries stable. With
		// DISTINCT projections, ordering remains under the caller's control.
		if !q.Unique {
			tx = tx.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}})
		}
		values := make([]*unitable.DbQuery, 0)
		if err := tx.Find(&values).Error; err != nil {
			return nil, err
		}
		return values, nil
	}
	values, err := load(qry)
	if err != nil {
		return nil, err
	}
	allQuery, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	allQuery.PageSize, allQuery.PageToken, allQuery.Skip = 0, "", 0
	all, err := load(allQuery)
	if err != nil {
		return nil, err
	}
	return &pb.ListDbQueriesResponse{DbQueries: values, TotalCount: int32(len(all)), NextPageToken: nextPageToken(qry, len(all))}, nil
}
