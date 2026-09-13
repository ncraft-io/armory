package handlers

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/xuri/excelize/v2"
)

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "unitable-tests-")
	if err != nil {
		panic(err)
	}
	path := filepath.Join(dir, "config.json")
	body := fmt.Sprintf(`{"db":{"driver":"sqlite","dsn":%q},"dataDB":{"default":"test","dbs":{"test":{"driver":"sqlite","dsn":%q},"other":{"driver":"sqlite","dsn":%q}}}}`, filepath.Join(dir, "meta.db"), filepath.Join(dir, "data.db"), filepath.Join(dir, "other.db"))
	if err := os.WriteFile(path, []byte(body), 0600); err != nil {
		panic(err)
	}
	if err := config.LoadFile(path); err != nil {
		panic(err)
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}
func requireOK(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
func testTable(t *testing.T, name string) (pb.UnitableServer, *unitable.Table) {
	t.Helper()
	s := NewService()
	table, err := s.CreateTable(context.Background(), &pb.CreateTableRequest{Database: "test", Table: &unitable.Table{Name: name, JsonStyle: "lowerCamel", Columns: []*unitable.Column{{Name: "id", Type: "string"}, {Name: "label", Type: "string"}, {Name: "item_count", Type: "integer"}, {Name: "enabled", Type: "bool"}}}})
	requireOK(t, err)
	return s, table
}
func TestSchemaLifecycle(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "schema_lifecycle")
	got, err := s.GetTable(ctx, &pb.GetTableRequest{Database: "test", Id: table.Id})
	requireOK(t, err)
	if len(got.Columns) != 4 {
		t.Fatalf("columns not loaded: %v", got)
	}
	col, err := s.CreateColumn(ctx, &pb.CreateColumnRequest{Database: "test", Table: table.Id, Column: &unitable.Column{Name: "extra", Type: "string", Show: true}})
	requireOK(t, err)
	_, err = s.UpdateColumn(ctx, &pb.UpdateColumnRequest{Database: "test", Table: table.Name, Id: col.Id, Column: &unitable.Column{Name: "renamed", Show: false}})
	requireOK(t, err)
	col, err = s.GetColumn(ctx, &pb.GetColumnRequest{Database: "test", Table: table.Name, Id: col.Id})
	requireOK(t, err)
	if col.Name != "renamed" || col.Show || col.Type != "string" {
		t.Fatalf("bad updated column: %v", col)
	}
	_, err = s.CreateRow(ctx, &pb.CreateRowRequest{Database: "test", Table: table.Name, Row: core.NewObject().SetString("id", "a").SetString("renamed", "present")})
	requireOK(t, err)
	_, err = s.BatchCreateColumns(ctx, &pb.BatchCreateColumnsRequest{Database: "test", Table: table.Name, Columns: []*unitable.Column{{Name: "second", Type: "integer"}, {Name: "third", Type: "string"}}})
	requireOK(t, err)
	listed, err := s.ListColumns(ctx, &pb.ListColumnsRequest{Database: "test", Table: table.Name, PageSize: 2})
	requireOK(t, err)
	if len(listed.Columns) != 2 || listed.TotalCount != 7 || listed.NextPageToken != "1" {
		t.Fatalf("bad pagination: %v", listed)
	}
	_, err = s.DeleteColumn(ctx, &pb.DeleteColumnRequest{Database: "test", Table: table.Name, Id: col.Id})
	requireOK(t, err)
	if synchro.GetDataDB("test").Migrator().HasColumn(table.Name, "renamed") {
		t.Fatal("physical column not deleted")
	}
	got, err = s.GetTable(ctx, &pb.GetTableRequest{Database: "test", Id: table.Id})
	requireOK(t, err)
	if len(got.Columns) != 6 {
		t.Fatalf("deleted column remains: %v", got)
	}
	_, err = s.DeleteTable(ctx, &pb.DeleteTableRequest{Database: "test", Id: table.Name, Force: true})
	requireOK(t, err)
	if synchro.GetDataDB("test").Migrator().HasTable(table.Name) {
		t.Fatal("physical table not deleted")
	}
	var count int64
	requireOK(t, model.GetColumnModel().DB.Model(&unitable.Column{}).Where("table_id = ?", table.Id).Count(&count).Error)
	if count != 0 {
		t.Fatal("orphan columns")
	}
}
func TestRowsAndPagination(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "rows_lifecycle")
	created, err := s.BatchCreateRows(ctx, &pb.BatchCreateRowsRequest{Database: "test", Table: table.Id, Rows: []*core.Object{core.NewObject().SetString("label", "one"), core.NewObject().SetString("label", "two")}})
	requireOK(t, err)
	if len(created.Objects) != 2 || created.Objects[0].GetString("id") == "" {
		t.Fatalf("missing generated IDs: %v", created)
	}
	id := created.Objects[0].GetString("id")
	updated, err := s.BatchUpdateRows(ctx, &pb.BatchUpdateRowsRequest{Database: "test", Table: table.Name, Rows: []*core.Object{core.NewObject().SetString("id", id).SetString("label", "").SetInt64("itemCount", 0).SetBool("enabled", false)}})
	requireOK(t, err)
	if len(updated.Objects) != 1 || updated.Objects[0].GetString("id") != id {
		t.Fatalf("wrong batch result: %v", updated)
	}
	row, err := s.GetRow(ctx, &pb.GetRowRequest{Database: "test", Table: table.Name, Id: id})
	requireOK(t, err)
	if row.GetString("label") != "" {
		t.Fatalf("zero value not persisted: %v", row)
	}
	_, err = s.UpdateRow(ctx, &pb.UpdateRowRequest{Database: "test", Table: table.Name, Id: id, Row: core.NewObject().SetString("id", "different").SetString("label", "bad")})
	if err == nil {
		t.Fatal("accepted mismatched ID")
	}
	page, err := s.ListRow(ctx, &pb.ListRowRequest{Database: "test", Table: table.Name, PageSize: 1})
	requireOK(t, err)
	if page.TotalCount != 2 || page.NextPageToken != "1" {
		t.Fatalf("bad first page: %v", page)
	}
	page, err = s.ListRow(ctx, &pb.ListRowRequest{Database: "test", Table: table.Name, PageSize: 1, PageToken: "1"})
	requireOK(t, err)
	if page.NextPageToken != "" {
		t.Fatalf("bad last page: %v", page)
	}
	rows, err := s.BatchGetRow(ctx, &pb.BatchGetRowRequest{Database: "test", Table: table.Name, Ids: []string{id}})
	requireOK(t, err)
	if len(rows.Objects) != 1 {
		t.Fatalf("bad batch get: %v", rows)
	}
	exported, err := s.ExportRow(ctx, &pb.ExportRowRequest{Database: "test", Table: table.Name, Filename: "data.xlsx"})
	requireOK(t, err)
	content := exported.Objects[0].GetValue("@fileContent").GetBytesVal()
	f, err := excelize.OpenReader(bytes.NewReader(content))
	requireOK(t, err)
	defer f.Close()
	cells, err := f.GetRows("Sheet1")
	requireOK(t, err)
	if len(cells) != 3 {
		t.Fatalf("bad export: %v", cells)
	}
	_, err = s.BatchDeleteRows(ctx, &pb.BatchDeleteRowsRequest{Database: "test", Table: table.Name, Ids: []string{id}})
	requireOK(t, err)
}
func TestBatchInsertRollback(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "rollback_rows")
	d := synchro.GetDataDB("test")
	requireOK(t, d.Exec(`CREATE TRIGGER reject_bad BEFORE INSERT ON rollback_rows WHEN NEW.label = 'bad' BEGIN SELECT RAISE(ABORT, 'rejected'); END`).Error)
	_, err := s.BatchCreateRows(ctx, &pb.BatchCreateRowsRequest{Database: "test", Table: table.Name, Rows: []*core.Object{core.NewObject().SetString("id", "first").SetString("label", "good"), core.NewObject().SetString("id", "second").SetString("label", "bad")}})
	if err == nil {
		t.Fatal("insert error swallowed")
	}
	var count int64
	requireOK(t, d.Table(table.Name).Count(&count).Error)
	if count != 0 {
		t.Fatalf("partial commit: %d", count)
	}
}
func TestSyncAndScoping(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	d := synchro.GetDataDB("test")
	requireOK(t, d.Exec(`CREATE TABLE synced (id TEXT PRIMARY KEY, score REAL, active BOOLEAN)`).Error)
	table, err := s.SyncTable(ctx, &pb.SyncTableRequest{Database: "test", Id: "synced"})
	requireOK(t, err)
	if len(table.Columns) != 3 {
		t.Fatalf("bad synced table %v", table)
	}
	second, err := s.SyncTable(ctx, &pb.SyncTableRequest{Database: "test", Id: "test.synced"})
	requireOK(t, err)
	if table.ColumnIndex()["score"].Id != second.ColumnIndex()["score"].Id {
		t.Fatal("sync changed column identity")
	}
	_, err = s.CreateTable(ctx, &pb.CreateTableRequest{Database: "other", Table: &unitable.Table{Name: "other_table", Columns: []*unitable.Column{{Name: "id", Type: "string"}}}})
	requireOK(t, err)
	list, err := s.ListTables(ctx, &pb.ListTablesRequest{Database: "other"})
	requireOK(t, err)
	if len(list.Tables) != 1 || list.Tables[0].Name != "other_table" {
		t.Fatalf("cross database leak: %v", list)
	}
	_, err = s.GetTable(ctx, &pb.GetTableRequest{Database: "other", Id: "test.synced"})
	if err == nil {
		t.Fatal("cross database ID accepted")
	}
	databases, err := s.ListDatabases(ctx, &pb.ListDatabasesRequest{PageSize: 1})
	requireOK(t, err)
	if databases.TotalCount != 2 || databases.NextPageToken != "1" {
		t.Fatalf("bad database list: %v", databases)
	}
}
func TestInputValidation(t *testing.T) {
	for _, in := range []interface{}{nil, (*pb.ListRowRequest)(nil), "string", &pb.ListRowRequest{PageSize: -1}, &pb.ListRowRequest{PageSize: 1, PageToken: "-1"}} {
		if _, err := ParseQuery(in); err == nil {
			t.Fatalf("accepted %#v", in)
		}
	}
	if NewUnitable() == nil || GetUnitable() == nil {
		t.Fatal("nil pool value")
	}
	if getColIndex(701) != "ZZ" || getColIndex(702) != "AAA" || getColIndex(16383) != "XFD" {
		t.Fatal("wrong excel column")
	}
}

func TestStatsAndQueryFailures(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "stats_rows")
	_, err := s.BatchCreateRows(ctx, &pb.BatchCreateRowsRequest{Database: "test", Table: table.Name, Rows: []*core.Object{core.NewObject().SetString("id", "one").SetString("label", "a").SetInt64("itemCount", 1), core.NewObject().SetString("id", "two").SetString("label", "a").SetInt64("itemCount", 2)}})
	requireOK(t, err)
	stats, err := s.GetRowStat(ctx, &pb.GetRowStatRequest{Database: "test", Table: table.Name, Stats: []string{"group label", "avg   item_count", "sum item_count"}})
	requireOK(t, err)
	values := stats.GetValue("item_count").GetValuesVal().GetVals()
	if len(values) != 1 || values[0].GetObjectVal().GetFloat64("avg") != 1.5 || values[0].GetObjectVal().GetInt64("sum") != 3 {
		t.Fatalf("wrong stats %v", stats)
	}
	for _, expr := range []string{"bad", "sum absent", "evil item_count", "sum label", "days item_count"} {
		if _, err := s.GetRowStat(ctx, &pb.GetRowStatRequest{Database: "test", Table: table.Name, Stats: []string{expr}}); err == nil {
			t.Fatalf("accepted %s", expr)
		}
	}
	col := table.ColumnIndex()["item_count"]
	_, err = s.UpdateColumn(ctx, &pb.UpdateColumnRequest{Database: "test", Table: table.Name, Id: col.Id, Column: &unitable.Column{Statistical: true}})
	requireOK(t, err)
	stats, err = s.GetRowStat(ctx, &pb.GetRowStatRequest{Database: "test", Table: table.Name})
	requireOK(t, err)
	if stats.GetValue("item_count") == nil {
		t.Fatalf("automatic stats absent: %v", stats)
	}
	// Database errors used to defer Close on a nil rows pointer.
	if _, err := s.ListRow(ctx, &pb.ListRowRequest{Database: "test", Table: table.Name, Filter: "absent == 3"}); err == nil {
		t.Fatal("expected query failure")
	}
}
func TestTimeStats(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	d := synchro.GetDataDB("test")
	requireOK(t, d.Exec(`CREATE TABLE timed_rows (id TEXT PRIMARY KEY, occurred_at DATETIME)`).Error)
	requireOK(t, d.Exec(`INSERT INTO timed_rows VALUES ('a', '2026-01-01 00:00:00'), ('b', '2026-01-02 00:00:00'), ('c', '2026-02-01 00:00:00')`).Error)
	_, err := s.SyncTable(ctx, &pb.SyncTableRequest{Database: "test", Id: "timed_rows"})
	requireOK(t, err)
	stats, err := s.GetRowStat(ctx, &pb.GetRowStatRequest{Database: "test", Table: "timed_rows", Stats: []string{"months occurred_at"}})
	requireOK(t, err)
	values := stats.GetValue("occurred_at").GetValuesVal().GetVals()
	if len(values) != 2 || values[0].GetObjectVal().GetString("months") != "2026-01" || values[0].GetObjectVal().GetInt64("count") != 2 {
		t.Fatalf("wrong time buckets: %v", stats)
	}
}
func TestColumnBatchValidationAndForcedTypeChange(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "column_batches")
	_, err := s.BatchCreateColumns(ctx, &pb.BatchCreateColumnsRequest{Database: "test", Table: table.Name, Columns: []*unitable.Column{{Name: "new_col", Type: "string"}, {Name: "label", Type: "string"}}})
	if err == nil {
		t.Fatal("duplicate accepted")
	}
	if synchro.GetDataDB("test").Migrator().HasColumn(table.Name, "new_col") {
		t.Fatal("batch partially migrated")
	}
	label, count := table.ColumnIndex()["label"], table.ColumnIndex()["item_count"]
	_, err = s.BatchUpdateColumn(ctx, &pb.BatchUpdateColumnRequest{Database: "test", Table: table.Name, Columns: []*unitable.Column{{Id: label.Id, DisplayName: "Label"}, {Id: count.Id, DisplayName: "Count"}}})
	requireOK(t, err)
	_, err = s.UpdateTable(ctx, &pb.UpdateTableRequest{Database: "test", Id: table.Name, Table: &unitable.Table{Columns: []*unitable.Column{{Id: label.Id, Name: "changed_label", Type: "integer"}}}})
	if err == nil {
		t.Fatal("unforced renamed type change accepted")
	}
	_, err = s.UpdateTable(ctx, &pb.UpdateTableRequest{Database: "test", Id: table.Id, Force: true, Table: &unitable.Table{Columns: []*unitable.Column{{Id: label.Id, Name: "changed_label", Type: "integer"}}}})
	requireOK(t, err)
	_, err = s.BatchDeleteColumn(ctx, &pb.BatchDeleteColumnRequest{Database: "test", Table: table.Name, Ids: []string{label.Id, count.Id}})
	requireOK(t, err)
	got, err := s.GetTable(ctx, &pb.GetTableRequest{Database: "test", Id: table.Name})
	requireOK(t, err)
	if len(got.Columns) != 2 {
		t.Fatalf("wrong remaining columns: %v", got)
	}
}
func TestIntegerRowIDs(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	_, err := s.CreateTable(ctx, &pb.CreateTableRequest{Database: "test", Table: &unitable.Table{Name: "integer_rows", Columns: []*unitable.Column{{Name: "id", Type: "integer"}, {Name: "label", Type: "string"}}}})
	requireOK(t, err)
	created, err := s.BatchCreateRows(ctx, &pb.BatchCreateRowsRequest{Database: "test", Table: "integer_rows", Rows: []*core.Object{core.NewObject().SetInt64("id", 42).SetString("label", "value")}})
	requireOK(t, err)
	if created.Objects[0].GetInt64("id") != 42 {
		t.Fatalf("integer ID lost: %v", created)
	}
	updated, err := s.BatchUpdateRows(ctx, &pb.BatchUpdateRowsRequest{Database: "test", Table: "integer_rows", Rows: []*core.Object{core.NewObject().SetInt64("id", 42).SetString("label", "")}})
	requireOK(t, err)
	if updated.Objects[0].GetInt64("id") != 42 {
		t.Fatalf("integer ID lost: %v", updated)
	}
	_, err = s.UpdateRow(ctx, &pb.UpdateRowRequest{Database: "test", Table: "integer_rows", Id: "42", Row: core.NewObject().SetString("label", "updated")})
	requireOK(t, err)
	row, err := s.GetRow(ctx, &pb.GetRowRequest{Database: "test", Table: "integer_rows", Id: "42"})
	requireOK(t, err)
	if row.GetString("label") != "updated" {
		t.Fatalf("numeric path ID update failed: %v", row)
	}

}

func TestQueryProjectionAndSkip(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "projection_rows")
	_, err := s.BatchCreateRows(ctx, &pb.BatchCreateRowsRequest{Database: "test", Table: table.Name, Rows: []*core.Object{core.NewObject().SetString("id", "a").SetInt64("itemCount", 1), core.NewObject().SetString("id", "b").SetInt64("itemCount", 2), core.NewObject().SetString("id", "c").SetInt64("itemCount", 3)}})
	requireOK(t, err)
	page, err := s.ListRow(ctx, &pb.ListRowRequest{Database: "test", Table: table.Name, PageSize: 1, Skip: 1, FieldMask: core.NewFieldMask("item_count")})
	requireOK(t, err)
	if page.TotalCount != 3 || page.NextPageToken != "1" || page.Objects[0].GetValue("itemCount") == nil {
		t.Fatalf("bad projection/count: %v", page)
	}
}

func TestDatabaseFilter(t *testing.T) {
	s := NewService()
	list, err := s.ListDatabases(context.Background(), &pb.ListDatabasesRequest{Filter: `name == "other"`, PageSize: 1})
	requireOK(t, err)
	if list.TotalCount != 1 || len(list.Databases) != 1 || list.Databases[0].Name != "other" {
		t.Fatalf("wrong catalog filter: %v", list)
	}
}

func TestUnknownDatabaseDoesNotUseDefault(t *testing.T) {
	if synchro.GetDataDB("missing") != nil {
		t.Fatal("unknown database falls back to default")
	}
	s := NewService()
	_, err := s.CreateTable(context.Background(), &pb.CreateTableRequest{Database: "missing", Table: &unitable.Table{Name: "must_not_exist", Columns: []*unitable.Column{{Name: "id", Type: "string"}}}})
	if err == nil {
		t.Fatal("created table in unknown database")
	}
	if synchro.GetDataDB("test").Migrator().HasTable("must_not_exist") {
		t.Fatal("wrote to default database")
	}
}
func TestClearUniqueColumnFlag(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "unique_flag")
	col, err := s.CreateColumn(ctx, &pb.CreateColumnRequest{Database: "test", Table: table.Name, Column: &unitable.Column{Name: "code", Type: "string", Unique: true}})
	requireOK(t, err)
	_, err = s.UpdateColumn(ctx, &pb.UpdateColumnRequest{Database: "test", Table: table.Name, Id: col.Id, Column: &unitable.Column{Unique: false}})
	requireOK(t, err)
	_, err = s.BatchCreateRows(ctx, &pb.BatchCreateRowsRequest{Database: "test", Table: table.Name, Rows: []*core.Object{core.NewObject().SetString("id", "one").SetString("code", "same"), core.NewObject().SetString("id", "two").SetString("code", "same")}})
	requireOK(t, err)
}

func TestRepeatedStringRows(t *testing.T) {
	ctx := context.Background()
	s, table := testTable(t, "array_rows")
	_, err := s.CreateColumn(ctx, &pb.CreateColumnRequest{Database: "test", Table: table.Name, Column: &unitable.Column{Name: "tags", Type: "string", Repeated: true}})
	requireOK(t, err)
	row, err := core.NewObjectFromKeyValues("id", "a", "tags", []string{"a,b", `a"b`, `a\b`, ""})
	requireOK(t, err)
	_, err = s.CreateRow(ctx, &pb.CreateRowRequest{Database: "test", Table: table.Name, Row: row})
	requireOK(t, err)
	got, err := s.GetRow(ctx, &pb.GetRowRequest{Database: "test", Table: table.Name, Id: "a"})
	requireOK(t, err)
	tags := got.GetStringArray("tags")
	if len(tags) != 4 || tags[0] != "a,b" || tags[1] != `a"b` || tags[2] != `a\b` || tags[3] != "" {
		t.Fatalf("array did not round trip: %v", got)
	}
	row = core.NewObject().SetString("id", "a").SetValue("tags", core.NewArrayValue())
	_, err = s.BatchUpdateRows(ctx, &pb.BatchUpdateRowsRequest{Database: "test", Table: table.Name, Rows: []*core.Object{row}})
	requireOK(t, err)
	got, err = s.GetRow(ctx, &pb.GetRowRequest{Database: "test", Table: table.Name, Id: "a"})
	requireOK(t, err)
	if len(got.GetStringArray("tags")) != 0 {
		t.Fatalf("array not cleared: %v", got)
	}
}
