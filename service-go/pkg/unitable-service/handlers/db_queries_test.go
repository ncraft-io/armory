package handlers

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/ncraft-io/armory/service-go/pkg/unitable-service/svc"
	nconfig "github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config/source/file"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"google.golang.org/protobuf/proto"
)

func scalarQuery(name string, value int) *unitable.DbQuery {
	return &unitable.DbQuery{Name: name, Sql: fmt.Sprintf("SELECT %d AS item_count", value), Columns: []*unitable.Column{{Name: "item_count", Type: "integer"}}}
}
func TestDbQueryLifecycleAndPriority(t *testing.T) {
	ctx := context.Background()
	s := unitableServer{instance: &Instance{Synchro: synchro.New(), Queries: map[string]*unitable.DbQuery{}}}
	input := scalarQuery("query-lifecycle", 1)
	created, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: input})
	requireOK(t, err)
	if input.Id != "" || input.Database != "" {
		t.Fatal("request mutated")
	}
	if created.Id != "test.query-lifecycle" || created.CreateTime == nil {
		t.Fatalf("bad identity: %v", created)
	}
	_, err = s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: scalarQuery(input.Name, 99)})
	if err == nil {
		t.Fatal("duplicate create overwritten")
	}
	assertValue := func(db string, n int64) {
		t.Helper()
		r, e := s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: db, Id: input.Name})
		requireOK(t, e)
		if len(r.Objects) != 1 || r.Objects[0].GetInt64("itemCount") != n || r.TotalCount != 1 {
			t.Fatalf("bad result: %v", r)
		}
	}
	assertValue("test", 1)
	_, err = s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: "other", Id: input.Name})
	if err == nil {
		t.Fatal("cross database query leak")
	}
	_, err = s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "other", Query: scalarQuery(input.Name, 2)})
	requireOK(t, err)
	assertValue("other", 2)
	global := scalarQuery(input.Name, 3)
	local := scalarQuery(input.Name, 4)
	local.Database = "test"
	s.instance.Queries[queryConfigKey(global)] = global
	s.instance.Queries[queryConfigKey(local)] = local
	assertValue("test", 4)
	assertValue("other", 3)
	stored, err := s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: created.Id})
	requireOK(t, err)
	if stored.Sql != input.Sql {
		t.Fatal("management API returns configuration instead of stored definition")
	}
	_, err = s.UpdateDbQuery(ctx, &pb.UpdateDbQueryRequest{Database: "test", Id: created.Id, Query: scalarQuery(input.Name, 5)})
	requireOK(t, err)
	assertValue("test", 4)
	delete(s.instance.Queries, queryConfigKey(local))
	assertValue("test", 3)
	delete(s.instance.Queries, queryConfigKey(global))
	assertValue("test", 5)
	bad := scalarQuery(input.Name, 6)
	bad.Sql = "SELECT missing FROM unknown_table"
	_, err = s.UpdateDbQuery(ctx, &pb.UpdateDbQueryRequest{Database: "test", Id: created.Id, Query: bad})
	if err == nil {
		t.Fatal("accepted invalid SQL update")
	}
	assertValue("test", 5)
	for _, db := range []string{"other", "missing"} {
		_, err = s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: db, Id: created.Id})
		if err == nil {
			t.Fatal("cross database identifier accepted")
		}
	}
	s.instance.Queries[queryConfigKey(global)] = global
	_, err = s.DeleteDbQuery(ctx, &pb.DeleteDbQueryRequest{Database: "test", Id: created.Id})
	requireOK(t, err)
	assertValue("test", 3)
	delete(s.instance.Queries, queryConfigKey(global))
	_, err = s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: "test", Id: input.Name})
	if err == nil {
		t.Fatal("deleted query remains cached")
	}
	assertValue("other", 2)
}

func TestDbQueryParametersAndLegacyListRow(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "db_query_rows")
	for i, label := range []string{"one", "two", "a'?[];--"} {
		_, err := s.CreateRow(ctx, &pb.CreateRowRequest{Database: "test", Table: table.Name, Row: core.NewObject().SetString("id", fmt.Sprint(i)).SetString("label", label).SetInt64("itemCount", int64(i+1))})
		requireOK(t, err)
	}
	q := &unitable.DbQuery{Name: "optional-query", Sql: "SELECT label, item_count FROM db_query_rows WHERE item_count >= ? [AND label IN ?] ORDER BY item_count", Parameters: []*unitable.DbQuery_Parameter{{Name: "minimum", Type: "integer"}, {Name: "labels", Type: "string", IsArray: true}}, Columns: []*unitable.Column{{Name: "label", Type: "string"}, {Name: "item_count", Type: "integer"}}}
	_, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: q})
	requireOK(t, err)
	r, err := s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: "test", Id: q.Name, Parameters: core.NewObject().SetInt64("minimum", 1), PageSize: 2})
	requireOK(t, err)
	if len(r.Objects) != 2 || r.TotalCount != 3 || r.NextPageToken != "1" {
		t.Fatalf("bad pagination: %v", r)
	}
	r, err = s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: "test", Id: q.Name, Parameters: core.NewObject().SetInt64("minimum", 1).SetString("labels", "a'?[];--")})
	requireOK(t, err)
	if len(r.Objects) != 1 || r.Objects[0].GetString("label") != "a'?[];--" {
		t.Fatalf("unsafe or incorrect argument binding: %v", r)
	}
	httpCtx := context.WithValue(ctx, "http-request-query", url.Values{"minimum": {"2"}, "labels": {"two", "a'?[];--"}})
	legacy, err := s.ListRow(httpCtx, &pb.ListRowRequest{Database: "test", Table: table.Name, Query: q.Name})
	requireOK(t, err)
	if legacy.TotalCount != 2 {
		t.Fatalf("legacy path cannot use stored queries: %v", legacy)
	}
	for _, params := range []*core.Object{nil, core.NewObject().SetString("minimum", "not-an-int")} {
		_, err = s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: "test", Id: q.Name, Parameters: params})
		if err == nil {
			t.Fatal("invalid parameters accepted")
		}
	}
	literal := &unitable.DbQuery{Name: "literal-query", Sql: "SELECT '?[keep]' AS label, ? AS item_count -- ? in comment\n", Parameters: []*unitable.DbQuery_Parameter{{Name: "n", Type: "integer"}}, Columns: q.Columns}
	_, err = s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: literal})
	requireOK(t, err)
	r, err = s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: "test", Id: literal.Name, Parameters: core.NewObject().SetInt64("n", 7)})
	requireOK(t, err)
	if r.Objects[0].GetString("label") != "?[keep]" || r.Objects[0].GetInt64("itemCount") != 7 {
		t.Fatalf("SQL literal changed: %v", r)
	}
}

func TestDbQueryValidation(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	cases := []struct {
		name   string
		change func(*unitable.DbQuery)
	}{
		{"empty", func(q *unitable.DbQuery) { q.Sql = "" }},
		{"write", func(q *unitable.DbQuery) { q.Sql = "DELETE FROM db_queries" }},
		{"multi", func(q *unitable.DbQuery) { q.Sql += "; SELECT 2" }},
		{"cte_write", func(q *unitable.DbQuery) {
			q.Sql = "WITH changed AS (DELETE FROM db_queries RETURNING *) SELECT * FROM changed"
		}},
		{"syntax", func(q *unitable.DbQuery) { q.Sql = "SELECT FROM" }},
		{"unknown_table", func(q *unitable.DbQuery) { q.Sql = "SELECT item_count FROM does_not_exist" }},
		{"missing_parameters", func(q *unitable.DbQuery) { q.Sql = "SELECT ? AS item_count" }},
		{"nil_parameter", func(q *unitable.DbQuery) { q.Parameters = append(q.Parameters, nil) }},
		{"bad_parameter_type", func(q *unitable.DbQuery) {
			q.Sql = "SELECT ? AS item_count"
			q.Parameters = []*unitable.DbQuery_Parameter{{Name: "x", Type: "blob"}}
		}},
		{"inconsistent_parameter", func(q *unitable.DbQuery) {
			q.Sql = "SELECT ? + ? AS item_count"
			q.Parameters = []*unitable.DbQuery_Parameter{{Name: "x", Type: "integer"}, {Name: "x", Type: "float"}}
		}},
		{"nil_column", func(q *unitable.DbQuery) { q.Columns = append(q.Columns, nil) }},
		{"duplicate_column", func(q *unitable.DbQuery) { q.Columns = append(q.Columns, q.Columns[0]) }},
		{"wrong_alias", func(q *unitable.DbQuery) { q.Columns[0].Name = "wrong" }},
		{"column_type", func(q *unitable.DbQuery) { q.Columns[0].Type = "object" }},
		{"wrong_database", func(q *unitable.DbQuery) { q.Database = "other" }},
		{"wrong_id", func(q *unitable.DbQuery) { q.Id = "other.query" }},
		{"invalid_optional_variant", func(q *unitable.DbQuery) {
			q.Sql = "SELECT 1 AS item_count [WHERE ?=1] AND 1=1"
			q.Parameters = []*unitable.DbQuery_Parameter{{Name: "n", Type: "integer"}}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			q := scalarQuery("invalid-"+tc.name, 1)
			tc.change(q)
			_, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: q})
			if err == nil {
				t.Fatal("accepted invalid definition")
			}
			var n int64
			requireOK(t, model.GetDbQueryModel().DB.Model(&unitable.DbQuery{}).Where("name = ?", q.Name).Count(&n).Error)
			if n != 0 {
				t.Fatal("invalid definition persisted")
			}
		})
	}
	q := scalarQuery("repeated-parameters", 1)
	q.Sql = "SELECT ? + ? AS item_count; -- ending\n"
	q.Parameters = []*unitable.DbQuery_Parameter{{Name: "n", Type: "integer"}, {Name: "n", Type: "integer"}}
	_, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: q})
	requireOK(t, err)
	r, err := s.RunDbQuery(ctx, &pb.RunDbQueryRequest{Database: "test", Id: q.Name, Parameters: core.NewObject().SetInt64("n", 3)})
	requireOK(t, err)
	if r.Objects[0].GetInt64("itemCount") != 6 {
		t.Fatal("repeated parameter failed")
	}
}
func TestQueryTemplateDialectBinding(t *testing.T) {
	q := &unitable.DbQuery{Sql: "SELECT $$?[literal]$$, array[1], '?' FROM things WHERE id IN (?) [AND label = ?]", Parameters: []*unitable.DbQuery_Parameter{{Name: "ids", Type: "integer", IsArray: true}, {Name: "label", Type: "string"}}}
	parsed, err := parseQueryTemplate(q.Sql)
	requireOK(t, err)
	sql, args, err := parsed.bind(q, map[string]interface{}{"ids": []int{1, 2}}, "postgres")
	requireOK(t, err)
	if !strings.Contains(sql, "id IN ($1,$2)") || !strings.Contains(sql, "array[1]") || strings.Contains(sql, "AND label") || !reflect.DeepEqual(args, []interface{}{int64(1), int64(2)}) {
		t.Fatalf("bad SQL binding: %s, %#v", sql, args)
	}
	q = proto.Clone(q).(*unitable.DbQuery)
	q.Sql = "SELECT ? AS a [WHERE ? = ?]"
	q.Parameters = []*unitable.DbQuery_Parameter{{Name: "base", Type: "integer"}, {Name: "a", Type: "integer"}, {Name: "b", Type: "integer"}}
	parsed, err = parseQueryTemplate(q.Sql)
	requireOK(t, err)
	_, _, err = parsed.bind(q, map[string]interface{}{"base": 1, "a": 1}, "sqlite")
	if err == nil {
		t.Fatal("partial optional condition accepted")
	}
}
func TestDbQueryListPaginationAndScope(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	for i := 0; i < 3; i++ {
		_, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: scalarQuery(fmt.Sprintf("paging-%d", i), i)})
		requireOK(t, err)
	}
	_, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "other", Query: scalarQuery("paging-0", 9)})
	requireOK(t, err)
	first, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Filter: `name >= "paging-0" and name <= "paging-2"`, PageSize: 2})
	requireOK(t, err)
	if len(first.DbQueries) != 2 || first.TotalCount != 3 || first.NextPageToken != "1" {
		t.Fatalf("bad list page: %v", first)
	}
	second, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Filter: `name >= "paging-0" and name <= "paging-2"`, PageSize: 2, PageToken: first.NextPageToken})
	requireOK(t, err)
	if len(second.DbQueries) != 1 || second.NextPageToken != "" {
		t.Fatalf("bad final list page: %v", second)
	}
}

func TestDbQueryHTTPRoutes(t *testing.T) {
	s := NewService()
	router := mux.NewRouter()
	endpoints := svc.Endpoints{CreateDbQueryEndpoint: svc.MakeCreateDbQueryEndpoint(s), UpdateDbQueryEndpoint: svc.MakeUpdateDbQueryEndpoint(s), GetDbQueryEndpoint: svc.MakeGetDbQueryEndpoint(s), ListDbQueriesEndpoint: svc.MakeListDbQueriesEndpoint(s), DeleteDbQueryEndpoint: svc.MakeDeleteDbQueryEndpoint(s), RunDbQueryEndpoint: svc.MakeRunDbQueryEndpoint(s)}
	svc.RegisterHttpHandler(router, endpoints, nil, log.NewNopLogger())
	base := "/armory/unitable/v1/databases/test/queries"
	request := func(method, path, body string) string {
		t.Helper()
		r := httptest.NewRequest(method, path, strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)
		if w.Code < 200 || w.Code >= 300 {
			t.Fatalf("%s %s: %d %s", method, path, w.Code, w.Body.String())
		}
		return w.Body.String()
	}
	body := `{"name":"http-query","sql":"SELECT ? AS item_count","parameters":["n:integer"],"columns":[{"name":"item_count","type":"integer"}]}`
	request("POST", base, body)
	got := request("POST", base+"/http-query:run", `{"n":42}`)
	if !strings.Contains(got, `"itemCount":42`) {
		t.Fatalf("HTTP body/path binding failed: %s", got)
	}
	request("GET", base+"/http-query", "")
	request("GET", base+"?pageSize=1", "")
	request("PUT", base+"/http-query", strings.Replace(body, "SELECT ? AS", "SELECT ? + 1 AS", 1))
	got = request("POST", base+"/http-query:run", `{"n":42}`)
	if !strings.Contains(got, `"itemCount":43`) {
		t.Fatalf("HTTP update failed: %s", got)
	}
	request("DELETE", base+"/http-query", "")
}

func TestExistingQueryConfigurationCompatibility(t *testing.T) {
	cfg, err := nconfig.NewConfig()
	requireOK(t, err)
	defer cfg.Close()
	requireOK(t, cfg.Load(file.NewSource(file.WithPath("../../../configs/queries.yaml"))))
	var config DBQueryConfig
	requireOK(t, cfg.Get("dbQuery").Scan(&config))
	if len(config.Queries) == 0 {
		t.Fatal("no configuration queries loaded")
	}
	for _, q := range config.Queries {
		t.Run(q.Name, func(t *testing.T) {
			requireOK(t, validateQueryColumns(q))
			if q.Sql == "" {
				return
			}
			parsed, err := parseQueryTemplate(q.Sql)
			requireOK(t, err)
			values := map[string]interface{}{}
			for _, p := range q.Parameters {
				if p == nil {
					t.Fatal("nil configuration parameter")
				}
				if p.IsArray {
					values[p.Name] = []string{"sample"}
				} else {
					values[p.Name] = "sample"
				}
			}
			_, _, err = parsed.bind(q, values, "postgres")
			requireOK(t, err)
			required := map[string]interface{}{}
			index := 0
			for _, part := range parsed.parts {
				for _, token := range part.tokens {
					if token.parameter {
						name := q.Parameters[index].Name
						index++
						if !part.optional {
							required[name] = values[name]
						}
					}
				}
			}
			_, _, err = parsed.bind(q, required, "postgres")
			requireOK(t, err)
		})
	}
}

func TestDbQueryConcurrentCreateDoesNotOverwrite(t *testing.T) {
	s := NewService()
	ctx := context.Background()
	type result struct {
		query *unitable.DbQuery
		err   error
	}
	results := make(chan result, 4)
	for i := 0; i < 4; i++ {
		go func(value int) {
			q, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: scalarQuery("concurrent-query", value)})
			results <- result{q, err}
		}(i)
	}
	var winner *unitable.DbQuery
	for i := 0; i < 4; i++ {
		r := <-results
		if r.err == nil {
			if winner != nil {
				t.Fatal("more than one create succeeded")
			}
			winner = r.query
		} else if !core.IsAlreadyExistsError(r.err) {
			t.Errorf("unexpected create error: %v", r.err)
		}
	}
	if winner == nil {
		t.Fatal("no query created")
	}
	stored, err := s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: winner.Id})
	requireOK(t, err)
	if stored.Sql != winner.Sql {
		t.Fatal("duplicate create overwrote the winning definition")
	}
}
