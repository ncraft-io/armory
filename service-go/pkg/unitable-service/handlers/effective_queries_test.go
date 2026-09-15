package handlers

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"github.com/ncraft-io/armory/service-go/pkg/unitable-service/svc"
	"google.golang.org/protobuf/proto"
)

func TestEffectiveDbQueryReads(t *testing.T) {
	ctx := context.Background()
	s := unitableServer{instance: &Instance{Synchro: synchro.New(), Queries: map[string]*unitable.DbQuery{}}}
	for _, name := range []string{"effective-a", "effective-c"} {
		_, err := s.CreateDbQuery(ctx, &pb.CreateDbQueryRequest{Database: "test", Query: scalarQuery(name, 1)})
		requireOK(t, err)
	}
	global := scalarQuery("effective-a", 2)
	local := scalarQuery("effective-a", 3)
	local.Database = "test"
	configOnly := scalarQuery("effective-b", 4)
	foreign := scalarQuery("effective-d", 5)
	foreign.Database = "other"
	for _, q := range []*unitable.DbQuery{global, local, configOnly, foreign} {
		s.instance.Queries[queryConfigKey(q)] = q
	}
	expected := proto.Clone(local).(*unitable.DbQuery)
	persisted, err := s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: local.Name})
	requireOK(t, err)
	if persisted.Sql != scalarQuery(local.Name, 1).Sql {
		t.Fatal("default Get no longer returns stored definition")
	}
	got, err := s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: "test." + local.Name, Effective: true})
	requireOK(t, err)
	if got.Sql != local.Sql || got.Id != "test.effective-a" || got.Database != "test" || got.JsonStyle != "lowerCamel" {
		t.Fatalf("wrong effective definition: %v", got)
	}
	got.Columns[0].Name = "changed"
	if !proto.Equal(local, expected) {
		t.Fatal("effective Get mutated configuration")
	}
	_, err = s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: configOnly.Name})
	if !core.IsNotFoundError(err) {
		t.Fatalf("configuration-only record leaked into default Get: %v", err)
	}
	got, err = s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: configOnly.Name, Effective: true})
	requireOK(t, err)
	if got.Sql != configOnly.Sql {
		t.Fatal("configuration-only query missing")
	}
	other, err := s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "other", Id: global.Name, Effective: true})
	requireOK(t, err)
	if other.Sql != global.Sql {
		t.Fatal("database-specific configuration leaked")
	}
	_, err = s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "other", Id: "test." + global.Name, Effective: true})
	if err == nil {
		t.Fatal("cross-database ID accepted")
	}
	_, err = s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: foreign.Name, Effective: true})
	if !core.IsNotFoundError(err) {
		t.Fatal("foreign-only configuration leaked")
	}
	filter := `name in ["effective-a", "effective-b", "effective-c", "effective-d"]`
	storedList, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Filter: filter})
	requireOK(t, err)
	if storedList.TotalCount != 2 {
		t.Fatalf("default List includes configuration: %v", storedList)
	}
	first, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: filter, PageSize: 2})
	requireOK(t, err)
	if first.TotalCount != 3 || first.NextPageToken != "1" || len(first.DbQueries) != 2 || first.DbQueries[0].Name != local.Name || first.DbQueries[0].Sql != local.Sql || first.DbQueries[1].Name != configOnly.Name {
		t.Fatalf("bad merged first page: %v", first)
	}
	second, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: filter, PageSize: 2, PageToken: "1"})
	requireOK(t, err)
	if second.TotalCount != 3 || second.NextPageToken != "" || len(second.DbQueries) != 1 || second.DbQueries[0].Name != "effective-c" {
		t.Fatalf("bad merged final page: %v", second)
	}
	for _, q := range first.DbQueries {
		single, e := s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: q.Id, Effective: true})
		requireOK(t, e)
		if !proto.Equal(q, single) {
			t.Fatalf("effective Get and List disagree: %v / %v", q, single)
		}
	}
	shadowed, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: `name == "effective-a" and sql == "SELECT 1 AS item_count"`})
	requireOK(t, err)
	if shadowed.TotalCount != 0 {
		t.Fatal("filter applied before override")
	}
	selected, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: `name == "effective-a" and sql == "SELECT 3 AS item_count"`})
	requireOK(t, err)
	if selected.TotalCount != 1 {
		t.Fatal("filter did not see effective SQL")
	}
	ordered, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: filter, Order: &core.Ordering{Orders: []*core.Ordering_Order{{Field: "name", Sort: core.Ordering_SORT_DESC}}}, Skip: 1, PageSize: 1})
	requireOK(t, err)
	if len(ordered.DbQueries) != 1 || ordered.DbQueries[0].Name != configOnly.Name {
		t.Fatalf("wrong ordering/skip: %v", ordered)
	}
	projected, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: filter, FieldMask: &core.FieldMask{Paths: []string{"name"}}})
	requireOK(t, err)
	if projected.TotalCount != 3 || projected.DbQueries[0].Name == "" || projected.DbQueries[0].Sql != "" {
		t.Fatalf("bad projection: %v", projected)
	}
	distinct, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: filter, Unique: true, FieldMask: &core.FieldMask{Paths: []string{"database"}}})
	requireOK(t, err)
	if distinct.TotalCount != 1 || len(distinct.DbQueries) != 1 || distinct.DbQueries[0].Database != "test" {
		t.Fatalf("bad merged distinct projection: %v", distinct)
	}
	empty, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: filter, PageSize: 2, PageToken: "5"})
	requireOK(t, err)
	if len(empty.DbQueries) != 0 || empty.TotalCount != 3 || empty.NextPageToken != "" {
		t.Fatalf("bad out-of-range page: %v", empty)
	}
	var count int64
	requireOK(t, model.GetDbQueryModel().DB.Model(&unitable.DbQuery{}).Where("name IN ?", []string{global.Name, configOnly.Name, "effective-c", foreign.Name}).Count(&count).Error)
	if count != 2 {
		t.Fatal("effective reads persisted configuration")
	}
	delete(s.instance.Queries, queryConfigKey(local))
	got, err = s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: global.Name, Effective: true})
	requireOK(t, err)
	if got.Sql != global.Sql {
		t.Fatal("global fallback failed")
	}
	delete(s.instance.Queries, queryConfigKey(global))
	got, err = s.GetDbQuery(ctx, &pb.GetDbQueryRequest{Database: "test", Id: global.Name, Effective: true})
	requireOK(t, err)
	if got.Sql != persisted.Sql {
		t.Fatal("stored fallback failed")
	}
}

func TestEffectiveQueryBindingsAndHTTP(t *testing.T) {
	ctx := context.Background()
	s := unitableServer{instance: &Instance{Synchro: synchro.New(), Queries: map[string]*unitable.DbQuery{}}}
	q := scalarQuery("effective-http", 7)
	q.Sql = "SELECT '?; ''literal'' --' AS item_count"
	q.Parameters = []*unitable.DbQuery_Parameter{{Name: "a", Type: "integer"}} // reads expose definitions without executing their SQL
	s.instance.Queries[q.Name] = q
	listed, err := s.ListDbQueries(ctx, &pb.ListDbQueriesRequest{Database: "test", Effective: true, Filter: `name == "effective-http"`})
	requireOK(t, err)
	if len(listed.DbQueries) != 1 || listed.DbQueries[0].Sql != q.Sql || listed.DbQueries[0].Parameters[0].Name != q.Parameters[0].Name {
		t.Fatalf("configuration bindings changed: %v", listed)
	}
	router := mux.NewRouter()
	svc.RegisterHttpHandler(router, svc.Endpoints{GetDbQueryEndpoint: svc.MakeGetDbQueryEndpoint(s), ListDbQueriesEndpoint: svc.MakeListDbQueriesEndpoint(s)}, nil, log.NewNopLogger())
	for _, path := range []string{"/effective-http?effective=true", "?effective=true"} {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, httptest.NewRequest("GET", "/armory/unitable/v1/databases/test/queries"+path, nil))
		if w.Code != 200 || !strings.Contains(w.Body.String(), "effective-http") {
			t.Fatalf("effective HTTP parameter not decoded: %d %s", w.Code, w.Body.String())
		}
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("GET", "/armory/unitable/v1/databases/test/queries/effective-http?effective=false", nil))
	if w.Code != 404 {
		t.Fatalf("false effective changed behavior: %d %s", w.Code, w.Body.String())
	}
}
