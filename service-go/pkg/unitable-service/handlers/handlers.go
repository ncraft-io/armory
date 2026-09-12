package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/hook"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/segmentio/ksuid"
	"github.com/xuri/excelize/v2"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/ncraft-io/armory/service-go/pkg/hook"

	// this service api
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
)

var (
	_ = unitable.Table{}
	_ = core.Null{}
	_ = unitable.Column{}
	_ = core.Object{}
)

type unitableServer struct {
	pb.UnimplementedUnitableServer
	instance *Instance
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.UnitableServer {
	unitableOnce.Do(func() {
		ut = &Instance{
			Synchro: synchro.New(),
			Queries: make(map[string]*unitable.DbQuery),
		}
		conf := &unitable.DbQueryConfig{}
		_ = config.ScanFrom(conf, "dbQuery")
		for _, dbQuery := range conf.Queries {
			ut.Queries[dbQuery.Name] = dbQuery
		}
	})

	return unitableServer{
		instance: ut,
	}
}

// CreateTable implements Interface.
func (s unitableServer) CreateTable(ctx context.Context, in *pb.CreateTableRequest) (*unitable.Table, error) {
	if in.Table == nil {
		return nil, core.NewInvalidArgumentError("not set the table body in request")
	}
	if len(in.Table.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database in request")
	}
	if len(in.Table.Name) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name in request")
	}
	if !nameRegex.MatchString(in.Table.Name) {
		return nil, core.NewInvalidArgumentError("the table name (%s) is invalid in request", in.Table.Name)
	}
	if len(in.Table.Columns) == 0 {
		return nil, core.NewInvalidArgumentError("the table has no columns in request")
	}
	for i, col := range in.Table.Columns {
		if len(col.Name) == 0 {
			return nil, core.NewInvalidArgumentError("the No. %d column in table has not set name", i)
		}
		if !col.IsTypeValid() {
			return nil, core.NewInvalidArgumentError("the No. %d column (%s) in table type %s is invalid", i, col.Name, col.Type)
		}
	}
	if len(in.Table.Id) == 0 {
		in.Table.Id = in.Table.Database + "." + in.Table.Name
	}
	if table, err := model.GetTableModel().Get(ctx, in.Table.Id); err == nil && table != nil {
		_, err = s.UpdateTable(ctx, &pb.UpdateTableRequest{Table: in.Table, Force: true})
		if err != nil {
			return nil, err
		}
		return &unitable.Table{Id: in.Table.Id}, nil
	}

	if in.Table.CreateTime == nil {
		in.Table.CreateTime = core.Now()
	}
	in.Table.UpdateTime = core.Now()
	if err := s.Synchro().MigrateTable(ctx, in.Table, nil, nil); err != nil {
		return nil, core.NewInternalError("failed to create table in %s", in.Table.Database)
	}

	for _, col := range in.Table.Columns {
		if len(col.Id) == 0 {
			col.Id = ksuid.New().String()
		}
		if col.CreateTime == nil {
			col.CreateTime = core.Now()
			col.UpdateTime = col.CreateTime
		} else {
			col.UpdateTime = core.Now()
		}
	}
	if _, err := model.GetTableModel().Create(ctx, in.Table); err != nil {
		return nil, err
	}

	resp := &unitable.Table{
		Id: in.Table.Id,
	}
	return resp, nil
}

func MergeColumn(target *unitable.Column, src *unitable.Column) {
	if len(target.Name) == 0 {
		target.Name = src.Name
	}
	if len(target.Type) == 0 {
		target.Type = src.Type
		target.Format = src.Format
	}
	if len(target.TableId) == 0 {
		target.TableId = src.TableId
	}
	if len(target.Database) == 0 {
		target.Database = src.Database
	}
	if len(target.DisplayName) == 0 {
		target.DisplayName = src.DisplayName
	}
	if len(target.ExportName) == 0 {
		target.ExportName = src.ExportName
	}
	if !target.Indexed {
		target.Indexed = src.Indexed
	}
}

// UpdateTable implements Interface.
func (s unitableServer) UpdateTable(ctx context.Context, in *pb.UpdateTableRequest) (*core.Null, error) {
	if in.Table == nil {
		return nil, core.NewInvalidArgumentError("not set the table body in request")
	}
	if len(in.Table.Database) == 0 {
		if len(in.Database) > 0 {
			in.Table.Database = in.Database
		} else {
			return nil, core.NewInvalidArgumentError("not set the database in request")
		}
	}
	if len(in.Table.Name) == 0 {
		if len(in.Id) > 0 {
			in.Table.Id = in.Id
			in.Table.Name = in.Id
		} else if len(in.Table.Id) > 0 {
			in.Table.Name = in.Table.Id
		} else {
			return nil, core.NewInvalidArgumentError("not set the table name or id in request")
		}
	}
	if !nameRegex.MatchString(in.Table.Name) {
		return nil, core.NewInvalidArgumentError("the table name (%s) is invalid in request", in.Table.Name)
	}
	if len(in.Table.Columns) == 0 {
		return nil, core.NewInvalidArgumentError("the table has no columns in request")
	}

	if len(in.Table.Id) == 0 {
		if len(in.Id) > 0 {
			in.Table.Id = in.Id
		} else {
			in.Table.Id = in.Table.Database + "." + in.Table.Name
		}
	}

	var dropCols []string
	renamedCols := make(map[string]string)
	if old, err := model.GetTableModel().Get(ctx, in.Table.Id); err != nil {
		return nil, core.NewInvalidArgumentError("the table %s not found, err: %s", in.Table.Id, err.Error())
	} else {
		columns := make(map[string]*unitable.Column)
		nameIndex := make(map[string]*unitable.Column)
		for _, col := range old.Columns {
			columns[col.Id] = col
			nameIndex[col.Name] = col
		}

		for _, col := range in.Table.Columns {
			if c, ok := nameIndex[col.Name]; ok {
				if len(col.Id) > 0 && col.Id != c.Id {
					return nil, core.NewInvalidArgumentError("the table %s update the column %s with different id, old: %s, new: %s", in.Table.Id, col.Name, c.Id, col.Id)
				}

				col.Id = c.Id
				MergeColumn(col, c)
				col.UpdateTime = core.Now()

				// NOT allow to change type if not force mode
				if col.Type != c.Type || (c.Type == "string" && c.Format != col.Format) {
					if in.Force {
						dropCols = append(dropCols, c.Name)
					} else {
						return nil, core.NewInvalidArgumentError("the table %s not to allow to change column %s type if not force mode", in.Table.Id, col.Name)
					}
				}
			} else if len(col.Id) == 0 {
				col.Id = ksuid.New().String()
				col.CreateTime = core.Now()
				col.UpdateTime = col.CreateTime
			} else {
				if c, ok := columns[col.Id]; ok {
					if len(col.Name) > 0 && len(c.Name) > 0 && col.Name != c.Name {
						renamedCols[c.Name] = col.Name
					}

					MergeColumn(col, c)
					col.UpdateTime = core.Now()
				} else {
					if col.CreateTime == nil {
						col.CreateTime = core.Now()
					}
					if col.UpdateTime == nil {
						col.UpdateTime = core.Now()
					}
				}
			}

			col.TableId = in.Table.Id
			if len(col.Type) == 0 {
				return nil, core.NewInvalidArgumentError("the column %s in table %s has not type set", col.Name, in.Table.Id)
			}
		}
	}

	in.Table.UpdateTime = core.Now()
	if err := s.Synchro().MigrateTable(ctx, in.Table, renamedCols, dropCols); err != nil {
		return nil, core.NewInternalError("failed to update table in %s, err: %s", in.Table.Database, err.Error())
	}

	columns := in.Table.Columns
	in.Table.Columns = nil
	if _, err := model.GetColumnModel().BatchCreate(ctx, columns...); err != nil {
		return nil, err
	}

	if _, err := model.GetTableModel().Update(ctx, in.Table); err != nil {
		return nil, err
	}

	return &core.Null{}, nil
}

// GetTable implements Interface.
func (s unitableServer) GetTable(ctx context.Context, in *pb.GetTableRequest) (*unitable.Table, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}

	id := in.Id
	if !strings.Contains(in.Id, ".") {
		id = in.Database + "." + in.Id
	}

	return model.GetTableModel().Get(ctx, id)
}

// ListTables implements Interface.
func (s unitableServer) ListTables(ctx context.Context, in *pb.ListTablesRequest) (*pb.ListTablesResponse, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}

	qry, err := ParseQuery(in)
	if err != nil {
		return nil, core.NewInvalidArgumentError("invalid query parameters, error: %s", err.Error())
	}
	if err = qry.Normalize(); err != nil {
		return nil, core.NewInvalidArgumentError("invalid query parameters, error: %s", err.Error())
	}

	tables, err := model.GetTableModel().List(ctx, qry)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListTablesResponse{
		Tables: tables,
	}
	return resp, nil
}

// DeleteTable implements Interface.
func (s unitableServer) DeleteTable(ctx context.Context, in *pb.DeleteTableRequest) (*core.Null, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}

	if _, err := model.GetTableModel().Delete(ctx, in.Id); err != nil {
		return nil, err
	}
	return &core.Null{}, nil
}

// SyncTable implements Interface.
func (s unitableServer) SyncTable(ctx context.Context, in *pb.SyncTableRequest) (*unitable.Table, error) {
	resp := &unitable.Table{
		// Id:
		// Name:
		// DisplayName:
		// ExportName:
		// Tenant:
		// Database:
		// Columns:
		// CreateTime:
		// UpdateTime:
		// DeleteTime:
	}
	return resp, nil
}

// CreateColumn implements Interface.
func (s unitableServer) CreateColumn(ctx context.Context, in *pb.CreateColumnRequest) (*unitable.Column, error) {
	resp := &unitable.Column{
		// Id:
		// Database:
		// Table:
		// Name:
		// DisplayName:
		// ExportName:
		// GroupDisplayName:
		// Type:
		// Format:
		// Indexed:
		// Unique:
		// Show:
		// Editable:
		// Filterable:
		// Temporal:
		// Dimensional:
		// Referenced:
		// CreateTime:
		// UpdateTime:
		// DeleteTime:
	}
	return resp, nil
}

// UpdateColumn implements Interface.
func (s unitableServer) UpdateColumn(ctx context.Context, in *pb.UpdateColumnRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// GetColumn implements Interface.
func (s unitableServer) GetColumn(ctx context.Context, in *pb.GetColumnRequest) (*unitable.Column, error) {
	resp := &unitable.Column{
		// Id:
		// Database:
		// Table:
		// Name:
		// DisplayName:
		// ExportName:
		// GroupDisplayName:
		// Type:
		// Format:
		// Indexed:
		// Unique:
		// Show:
		// Editable:
		// Filterable:
		// Temporal:
		// Dimensional:
		// Referenced:
		// CreateTime:
		// UpdateTime:
		// DeleteTime:
	}
	return resp, nil
}

// DeleteColumn implements Interface.
func (s unitableServer) DeleteColumn(ctx context.Context, in *pb.DeleteColumnRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// ListColumns implements Interface.
func (s unitableServer) ListColumns(ctx context.Context, in *pb.ListColumnsRequest) (*pb.ListColumnsResponse, error) {
	resp := &pb.ListColumnsResponse{
		// Columns:
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}

// BatchCreateColumns implements Interface.
func (s unitableServer) BatchCreateColumns(ctx context.Context, in *pb.BatchCreateColumnsRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// BatchUpdateColumn implements Interface.
func (s unitableServer) BatchUpdateColumn(ctx context.Context, in *pb.BatchUpdateColumnRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// BatchDeleteColumn implements Interface.
func (s unitableServer) BatchDeleteColumn(ctx context.Context, in *pb.BatchDeleteColumnRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// CreateRow implements Interface.
func (s unitableServer) CreateRow(ctx context.Context, in *pb.CreateRowRequest) (*core.Object, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if in.Row == nil || len(in.Row.Vals) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	id := in.Row.GetString("id")
	if len(id) == 0 {
		id = ksuid.New().String()
		in.Row.SetString("id", id)
	}

	tableID := in.Database + "." + in.Table
	if _, err := s.Synchro().InsertRow(ctx, tableID, in.Row); err != nil {
		return nil, core.NewInternalError("failed to insert the row in %s, (%v)", in.Table, in.Row.ToMapInterface())
	}

	hook.GetHook().Run(ctx)

	resp := &core.Object{}
	resp.SetString("id", id)
	return resp, nil
}

// UpdateRow implements Interface.
func (s unitableServer) UpdateRow(ctx context.Context, in *pb.UpdateRowRequest) (*core.Null, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row id")
	}
	if in.Row == nil || len(in.Row.Vals) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	id := in.Row.GetString("id")
	if len(id) == 0 {
		in.Row.SetString("id", in.Id)
	}

	tableId := in.Database + "." + in.Table
	if _, err := s.Synchro().UpdateRow(ctx, tableId, in.Row); err != nil {
		return nil, core.NewInternalError("failed to update the row in %s, (%v), err: %s", in.Table, in.Row.ToMapInterface(), err.Error())
	}

	hook.GetHook().Run(ctx)
	return &core.Null{}, nil
}

// GetRow implements Interface.
func (s unitableServer) GetRow(ctx context.Context, in *pb.GetRowRequest) (*core.Object, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row id")
	}

	tableId := in.Database + "." + in.Table
	if row, err := s.Synchro().GetRow(ctx, tableId, in.Id); err != nil {
		return nil, core.NewInternalError("failed to get the row %s in %s, err: %s", in.Id, in.Table, err.Error())
	} else {
		return row, nil
	}
}

// BatchGetRow implements Interface.
func (s unitableServer) BatchGetRow(ctx context.Context, in *pb.BatchGetRowRequest) (*pb.BatchGetRowResponse, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Ids) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row id")
	}
	var ids []string
	for _, id := range in.Ids {
		if len(id) > 0 {
			ids = append(ids, id)
		}
	}
	if len(ids) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row id")
	}

	tableId := in.Database + "." + in.Table
	if rows, err := s.Synchro().BatchGetRow(ctx, tableId, ids...); err != nil {
		return nil, core.NewInternalError("failed to get the row %s in %s, err: %s", ids, in.Table, err.Error())
	} else {
		return &pb.BatchGetRowResponse{
			Objects: rows,
		}, nil
	}
}

// DeleteRow implements Interface.
func (s unitableServer) DeleteRow(ctx context.Context, in *pb.DeleteRowRequest) (*core.Null, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row id")
	}

	tableId := in.Database + "." + in.Table
	if _, err := s.Synchro().DeleteRows(ctx, tableId, in.Id); err != nil {
		return nil, core.NewInternalError("failed to delete the row %s in %s, err: %s", in.Id, in.Table, err.Error())
	} else {
		hook.GetHook().Run(ctx)
		return &core.Null{}, nil
	}
}

var optStr = regexp.MustCompile(`\[[^\[\]]+]`)

func isEmpty(values []interface{}) bool {
	if len(values) == 0 {
		return true
	}
	allNil := true
	for _, v := range values {
		if v != nil {
			allNil = false
		}
	}
	return allNil
}

// ListRow implements Interface.
func (s unitableServer) ListRow(ctx context.Context, in *pb.ListRowRequest) (*pb.ListRowResponse, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}

	if q, ok := s.Queries()[in.Query]; ok && len(in.Query) > 0 {
		var values []interface{}
		query := &unitable.DbQuery{
			Id:         q.Id,
			Name:       q.Name,
			Sql:        q.Sql,
			Parameters: q.Parameters,
			Columns:    q.Columns,
			JsonStyle:  q.JsonStyle,
		}
		if len(query.JsonStyle) == 0 {
			query.JsonStyle = synchro.LowerCamel
		}

		if vals, ok := ctx.Value("http-request-query").(url.Values); ok {
			for _, p := range query.Parameters {
				if v, ok := vals[p.Name]; ok && len(v) > 0 {
					if p.IsArray {
						if len(v) == 1 {
							v = strings.Split(v[0], ",")
						}

						if p.PgArray {
							//values = append(values, pgtype.FlatArray[string](v))
							values = append(values, core.NewStringValues(v...))
						} else {
							values = append(values, v)
						}
					} else {
						values = append(values, v[0])
					}
				} else {
					values = append(values, nil)
				}
			}
		}

		if len(query.Sql) == 0 {
			example := query.Example()
			return &pb.ListRowResponse{
				Objects:    []*core.Object{example},
				TotalCount: 1,
			}, nil
		} else {
			if isEmpty(values) {
				query.Sql = string(optStr.ReplaceAll([]byte(query.Sql), []byte("")))
				values = []interface{}{}
			} else {
				query.Sql = strings.Replace(query.Sql, "[", "", -1)
				query.Sql = strings.Replace(query.Sql, "]", "", -1)
			}

			objs, err := s.Synchro().QueryBy(ctx, in.Database, in.Table, query, values)
			if err != nil {
				return nil, err
			}
			return &pb.ListRowResponse{
				Objects:    objs,
				TotalCount: 1,
			}, nil
		}
	}

	qry, err := ParseQuery(in)
	if err != nil {
		return nil, core.NewInvalidArgumentError("invalid query parameters, error: %s", err.Error())
	}
	if err = qry.Normalize(); err != nil {
		return nil, core.NewInvalidArgumentError("invalid query parameters, error: %s", err.Error())
	}

	tableId := in.Database + "." + in.Table
	if rows, totalCnt, err := s.Synchro().QueryRows(ctx, tableId, qry); err != nil {
		return nil, core.NewInternalError("failed to query the row in %s, err: %s", tableId, err.Error())
	} else {
		index := ""
		if len(in.PageToken) > 0 {
			t, _ := strconv.ParseInt(in.PageToken, 10, 64)
			t += 1
			if t*int64(in.PageSize) < int64(totalCnt) {
				index = fmt.Sprint(t)
			}
		}

		return &pb.ListRowResponse{
			Objects:       rows,
			TotalCount:    int32(totalCnt),
			NextPageToken: index,
		}, nil
	}
}

// ExportRow implements Interface.
func (s unitableServer) ExportRow(ctx context.Context, in *pb.ExportRowRequest) (*pb.ExportRowResponse, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}

	qry, err := ParseQuery(in)
	if err != nil {
		return nil, core.NewInvalidArgumentError("invalid query parameters, error: %s", err.Error())
	}
	if err = qry.Normalize(); err != nil {
		return nil, core.NewInvalidArgumentError("invalid query parameters, error: %s", err.Error())
	}

	tableId := in.Database + "." + in.Table
	if rows, totalCnt, err := s.Synchro().QueryRows(ctx, tableId, qry); err != nil {
		return nil, core.NewInternalError("failed to query the row in %s, err: %s", tableId, err.Error())
	} else {
		if len(in.Filename) > 0 {
			meta := s.Synchro().GetMetaTable(synchro.TableId(in.Database, in.Table), nil)
			columns := meta.Table.Columns

			f := excelize.NewFile()
			defer func() {
				if err = f.Close(); err != nil {
					logs.Warnw("failed to close the excelize's file", "error", err)
				}
			}()

			sname := "Sheet1"
			_, err := f.NewSheet(sname)
			if err != nil {
				return nil, err
			}

			colIndex := make(map[string]int)
			for i, col := range columns {
				_ = f.SetCellValue(sname, getColIndex(i)+"1", col.DisplayName)
				colIndex[col.Name] = i
			}

			for i, row := range rows {
				vals := row.GetVals()
				for k, v := range vals {
					col := colIndex[strcase.ToSnake(k)]
					cell := getColIndex(col) + strconv.Itoa(i+2)
					switch v.GetKind() {
					case core.ValueKind_VALUE_KIND_NULL:
					case core.ValueKind_VALUE_KIND_BOOLEAN:
						_ = f.SetCellValue(sname, cell, v.GetBoolVal())
					case core.ValueKind_VALUE_KIND_INTEGER:
						_ = f.SetCellValue(sname, cell, v.GetInt64())
					case core.ValueKind_VALUE_KIND_NUMBER:
						_ = f.SetCellValue(sname, cell, v.GetFloat64())
					case core.ValueKind_VALUE_KIND_STRING:
						_ = f.SetCellValue(sname, cell, v.GetStringVal())
					}
				}
			}
			buf := bytes.NewBuffer(nil)
			if err = f.Write(buf); err != nil {
				return nil, err
			}

			return &pb.ExportRowResponse{
				Objects: []*core.Object{core.NewObject().
					SetString("@fileName", in.Filename).
					SetValue("@fileContent", core.NewBytesValue(buf.Bytes()))},
			}, nil
		} else {
			return &pb.ExportRowResponse{
				Objects:    rows,
				TotalCount: int32(totalCnt),
			}, nil
		}
	}
}

// BatchCreateRows implements Interface.
func (s unitableServer) BatchCreateRows(ctx context.Context, in *pb.BatchCreateRowsRequest) (*pb.BatchCreateRowsResponse, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Rows) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	for _, row := range in.Rows {
		idv := row.GetValue("id")
		if idv != nil {
			if idv.GetInt64() == 0 {
				id := idv.GetString()
				if len(id) == 0 {
					id = ksuid.New().String()
					row.SetString("id", id)
				}
			}
		}
	}

	resp := &pb.BatchCreateRowsResponse{}
	tableId := in.Database + "." + in.Table
	if _, irs, err := s.Synchro().InsertRows(ctx, tableId, in.Rows...); err != nil {
		logs.ErrLogw("failed to batch create the rows", "table", tableId, "error", err)
		return nil, core.NewInternalError("failed to batch create the rows in %s, err: %s", tableId, err.Error())
	} else {
		for i, ir := range irs {
			if ir > 0 {
				id := in.Rows[i].GetString("id")
				resp.Objects = append(resp.Objects, core.NewObject().SetString("id", id))
			} else {
				resp.Objects = append(resp.Objects, nil)
			}
		}
	}

	return resp, nil
}

// BatchUpdateRows implements Interface.
func (s unitableServer) BatchUpdateRows(ctx context.Context, in *pb.BatchUpdateRowsRequest) (*pb.BatchUpdateRowsResponse, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Rows) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	for i, row := range in.Rows {
		idv := row.GetValue("id")
		if idv == nil || idv.GetInt64() == 0 || len(idv.GetString()) == 0 {
			return nil, core.NewInvalidArgumentError("the No. %d (begin with 1) row have not set the id in batch", i+1)
		}
	}

	resp := &pb.BatchUpdateRowsResponse{}
	tableId := in.Database + "." + in.Table
	if _, irs, err := s.Synchro().UpdateInsertRows(ctx, tableId, in.Rows...); err != nil {
		logs.ErrLogw("failed to batch update the rows ", "table", tableId)
		return nil, core.NewInternalError("failed to batch update the rows in %s, err: %s", tableId, err.Error())
	} else {
		for _, ir := range irs {
			if ir > 0 {
				id := in.Rows[ir].GetString("id")
				resp.Objects = append(resp.Objects, core.NewObject().SetString("id", id))
			} else {
				resp.Objects = append(resp.Objects, nil)
			}
		}
	}

	return resp, nil
}

// BatchDeleteRows implements Interface.
func (s unitableServer) BatchDeleteRows(ctx context.Context, in *pb.BatchDeleteRowsRequest) (*core.Null, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Ids) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row ids")
	}

	tableId := in.Database + "." + in.Table
	if _, err := s.Synchro().DeleteRows(ctx, tableId, in.Ids...); err != nil {
		return nil, core.NewInternalError("failed to batch delete the row in %s, err: %s", tableId, err.Error())
	}

	return &core.Null{}, nil
}

// ListDatabases implements Interface.
func (s unitableServer) ListDatabases(ctx context.Context, in *pb.ListDatabasesRequest) (*pb.ListDatabasesResponse, error) {
	resp := &pb.ListDatabasesResponse{
		// Databases:
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}

// GetRowStat implements Interface.
func (s unitableServer) GetRowStat(ctx context.Context, in *pb.GetRowStatRequest) (*core.Object, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	if len(in.Stats) == 0 {
		return nil, core.NewInvalidArgumentError("not set the stat expression")
	}

	stat := strings.Join(in.Stats, "|")

	qry, err := ParseQuery(in)
	if err != nil {
		return nil, core.NewInvalidArgumentError("invalid the filter expression: %s", in.Filter)
	}

	exprs := strings.Split(stat, "|")
	cals := make(map[string][]string)
	for _, expr := range exprs {
		expr = strings.TrimSpace(expr)
		segments := strings.Split(expr, " ")
		if len(segments) == 2 {
			fun := strings.TrimSpace(segments[0])
			op := strings.TrimSpace(segments[1])
			if fun == "group" {
				qry.Groups = append(qry.Groups, op)
			} else {
				cals[op] = append(cals[op], fun)
			}
		} else if len(segments) == 3 {
		} else if len(segments) == 4 {
		}
	}

	for op, funs := range cals {
		prj := &query.FieldProjection{
			Name:      op,
			Functions: funs,
			Alias:     nil,
		}

		qry.Projections = append(qry.Projections, prj)
	}

	if err = qry.Normalize(); err != nil {
		return nil, core.NewInvalidArgumentError("invalid query or stats expressions, error: %s", err.Error())
	}

	tableId := in.Database + "." + in.Table
	rows, err := s.Synchro().CalcStats(ctx, tableId, qry)
	if err != nil {
		return nil, core.NewInternalError("failed to get data from db, error: %s", err.Error())
	}

	if len(rows) == 0 {
		return nil, core.NewNotFoundError("failed to found data from db")
	} else if len(rows) > 0 {
		resp := &core.Object{
			Vals: make(map[string]*core.Value),
		}

		groups := make(map[string]bool)
		for _, group := range qry.Groups {
			groups[group] = true
		}

		fieldNames := make(map[string]bool)
		for k := range rows[0].GetVals() {
			if ok := groups[k]; ok {
				continue
			}

			name := qry.GetField(k).GetName()
			if len(name) > 0 {
				fieldNames[name] = true
			}
		}

		for field := range fieldNames {
			var nrs []*core.Value
			for _, row := range rows {
				nr := core.NewObject()
				for k, v := range row.GetVals() {
					if ok := groups[k]; ok {
						nr.SetValue(k, v)
						continue
					}

					f := qry.GetField(k)
					name, fun := f.GetName(), f.GetFunction()
					if name == field {
						nr.SetValue(fun, v)
					}
				}

				nrs = append(nrs, core.NewObjectValue(nr))
			}
			resp.SetValue(field, core.NewArrayValue(nrs...))
		}
		return resp, nil
	}
	return nil, nil
}
