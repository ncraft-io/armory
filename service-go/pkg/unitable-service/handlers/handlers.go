package handlers

import (
	"bytes"
	"context"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/hook"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/segmentio/ksuid"
	"github.com/xuri/excelize/v2"
	"google.golang.org/protobuf/proto"
	"net/url"
	"strconv"

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
			if dbQuery != nil {
				ut.Queries[queryConfigKey(dbQuery)] = dbQuery
			}
		}
	})

	return unitableServer{
		instance: ut,
	}
}

// CreateRow implements Interface.
func (s unitableServer) CreateRow(ctx context.Context, in *pb.CreateRowRequest) (*core.Object, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.CreateRowRequest)
	in.Table = tableName
	if in.Row == nil || len(in.Row.Vals) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	if !validRowID(in.Row.GetValue("id")) {
		in.Row.SetString("id", ksuid.New().String())
	}

	tableID := in.Database + "." + in.Table
	if _, err := s.Synchro().InsertRow(ctx, tableID, in.Row); err != nil {
		return nil, core.NewInternalError("failed to insert the row in %s, (%v)", in.Table, in.Row.ToMapInterface())
	}

	hook.GetHook().Run(ctx)

	resp := &core.Object{}
	resp.SetValue("id", in.Row.GetValue("id"))
	return resp, nil
}

// UpdateRow implements Interface.
func (s unitableServer) UpdateRow(ctx context.Context, in *pb.UpdateRowRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.UpdateRowRequest)
	in.Table = tableName
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row id")
	}
	if in.Row == nil || len(in.Row.Vals) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	id := in.Row.GetValue("id")
	if validRowID(id) && rowIDString(id) != in.Id {
		return nil, core.NewInvalidArgumentError("row id differs from request path")
	}
	if !validRowID(id) {
		meta := s.Synchro().GetMetaTable(in.Database+"."+in.Table, nil)
		if meta == nil {
			return nil, core.NewNotFoundError("table not found")
		}
		if col := meta.Table.ColumnIndex()["id"]; col != nil && col.Type == "integer" {
			value, err := strconv.ParseInt(in.Id, 10, 64)
			if err != nil || value == 0 {
				return nil, core.NewInvalidArgumentError("invalid integer row id")
			}
			in.Row.SetInt64("id", value)
		} else {
			in.Row.SetString("id", in.Id)
		}
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
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.GetRowRequest)
	in.Table = tableName
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
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.BatchGetRowRequest)
	in.Table = tableName
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
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.DeleteRowRequest)
	in.Table = tableName
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

// ListRow implements Interface.
func (s unitableServer) ListRow(ctx context.Context, in *pb.ListRowRequest) (*pb.ListRowResponse, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.ListRowRequest)
	in.Table = tableName

	if in.Query != "" {
		values := map[string]interface{}{}
		if vals, ok := ctx.Value("http-request-query").(url.Values); ok {
			for key, vals := range vals {
				if len(vals) == 1 {
					values[key] = vals[0]
				} else if len(vals) > 1 {
					values[key] = vals
				}
			}
		}
		objects, err := s.runQuery(ctx, in.Database, in.Query, values)
		if err != nil {
			return nil, err
		}
		return &pb.ListRowResponse{Objects: objects, TotalCount: int32(len(objects))}, nil
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

		return &pb.ListRowResponse{
			Objects:       rows,
			TotalCount:    int32(totalCnt),
			NextPageToken: nextPageToken(qry, totalCnt),
		}, nil
	}
}

// ExportRow implements Interface.
func (s unitableServer) ExportRow(ctx context.Context, in *pb.ExportRowRequest) (*pb.ExportRowResponse, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.ExportRowRequest)
	in.Table = tableName

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
			if meta == nil || meta.Table == nil {
				return nil, core.NewNotFoundError("table metadata not found")
			}
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
				label := col.ExportName
				if label == "" {
					label = col.DisplayName
				}
				if label == "" {
					label = col.Name
				}
				if err := f.SetCellValue(sname, getColIndex(i)+"1", label); err != nil {
					return nil, err
				}
				colIndex[col.Name] = i
			}

			for i, row := range rows {
				vals := row.GetVals()
				for k, v := range vals {
					col, ok := colIndex[strcase.ToSnake(k)]
					if !ok {
						continue
					}
					cell := getColIndex(col) + strconv.Itoa(i+2)
					switch v.GetKind() {
					case core.ValueKind_VALUE_KIND_NULL:
					case core.ValueKind_VALUE_KIND_BOOLEAN:
						if err := f.SetCellValue(sname, cell, v.GetBoolVal()); err != nil {
							return nil, err
						}
					case core.ValueKind_VALUE_KIND_INTEGER:
						if err := f.SetCellValue(sname, cell, v.GetInt64()); err != nil {
							return nil, err
						}
					case core.ValueKind_VALUE_KIND_NUMBER:
						if err := f.SetCellValue(sname, cell, v.GetFloat64()); err != nil {
							return nil, err
						}
					case core.ValueKind_VALUE_KIND_STRING:
						if err := f.SetCellValue(sname, cell, v.GetStringVal()); err != nil {
							return nil, err
						}
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
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.BatchCreateRowsRequest)
	in.Table = tableName
	if len(in.Rows) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	for i, row := range in.Rows {
		if row == nil || len(row.Vals) == 0 {
			return nil, core.NewInvalidArgumentError("empty row at index %d", i)
		}
		if !validRowID(row.GetValue("id")) {
			row.SetString("id", ksuid.New().String())
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
				resp.Objects = append(resp.Objects, core.NewObject().SetValue("id", in.Rows[i].GetValue("id")))
			} else {
				resp.Objects = append(resp.Objects, nil)
			}
		}
	}

	hook.GetHook().Run(ctx)
	return resp, nil
}

// BatchUpdateRows implements Interface.
func (s unitableServer) BatchUpdateRows(ctx context.Context, in *pb.BatchUpdateRowsRequest) (*pb.BatchUpdateRowsResponse, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.BatchUpdateRowsRequest)
	in.Table = tableName
	if len(in.Rows) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row")
	}

	for i, row := range in.Rows {
		idv := row.GetValue("id")
		if !validRowID(idv) {
			return nil, core.NewInvalidArgumentError("the No. %d (begin with 1) row have not set the id in batch", i+1)
		}
	}

	resp := &pb.BatchUpdateRowsResponse{}
	tableId := in.Database + "." + in.Table
	if _, irs, err := s.Synchro().UpdateInsertRows(ctx, tableId, in.Rows...); err != nil {
		logs.ErrLogw("failed to batch update the rows ", "table", tableId)
		return nil, core.NewInternalError("failed to batch update the rows in %s, err: %s", tableId, err.Error())
	} else {
		for i, ir := range irs {
			if ir > 0 {
				resp.Objects = append(resp.Objects, core.NewObject().SetValue("id", in.Rows[i].GetValue("id")))
			} else {
				resp.Objects = append(resp.Objects, nil)
			}
		}
	}

	hook.GetHook().Run(ctx)
	return resp, nil
}

// BatchDeleteRows implements Interface.
func (s unitableServer) BatchDeleteRows(ctx context.Context, in *pb.BatchDeleteRowsRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}
	if len(in.Table) == 0 {
		return nil, core.NewInvalidArgumentError("not set the table name")
	}
	_, tableName, identityErr := tableIdentity(in.Database, in.Table)
	if identityErr != nil {
		return nil, identityErr
	}
	in = proto.Clone(in).(*pb.BatchDeleteRowsRequest)
	in.Table = tableName
	if len(in.Ids) == 0 {
		return nil, core.NewInvalidArgumentError("not set the row ids")
	}

	tableId := in.Database + "." + in.Table
	if _, err := s.Synchro().DeleteRows(ctx, tableId, in.Ids...); err != nil {
		return nil, core.NewInternalError("failed to batch delete the row in %s, err: %s", tableId, err.Error())
	}

	hook.GetHook().Run(ctx)
	return &core.Null{}, nil
}

func validRowID(id *core.Value) bool {
	if id == nil {
		return false
	}
	switch id.GetKind() {
	case core.ValueKind_VALUE_KIND_STRING:
		return id.GetStringVal() != ""
	case core.ValueKind_VALUE_KIND_INTEGER:
		return id.GetInt64() != 0
	default:
		return false
	}
}
func rowIDString(id *core.Value) string {
	if id.GetKind() == core.ValueKind_VALUE_KIND_INTEGER {
		return strconv.FormatInt(id.GetInt64(), 10)
	}
	return id.GetStringVal()
}
