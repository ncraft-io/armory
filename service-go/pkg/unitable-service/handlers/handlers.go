package handlers

import (
	"context"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"regexp"

	"github.com/segmentio/ksuid"

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

	Synchro *synchro.Synchro
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.UnitableServer {
	return unitableServer{
		Synchro: synchro.New(),
	}
}

var nameRegex = regexp.MustCompile(`[a-z][a-z0-9]*`)

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

	in.Table.Id = in.Table.Name
	in.Table.CreateTime = core.Now()
	in.Table.UpdateTime = core.Now()
	if err := s.Synchro.CreateTable(ctx, in.Table); err != nil {
		return nil, core.NewInternalError("failed to create table in %s", in.Table.Database)
	}

	for _, col := range in.Table.Columns {
		if len(col.Id) == 0 {
			col.Id = ksuid.New().String()
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

// UpdateTable implements Interface.
func (s unitableServer) UpdateTable(ctx context.Context, in *pb.UpdateTableRequest) (*core.Null, error) {
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

	in.Table.Id = in.Table.Name
	in.Table.UpdateTime = core.Now()
	if err := s.Synchro.MigrateTable(ctx, in.Table); err != nil {
		return nil, core.NewInternalError("failed to update table in %s", in.Table.Database)
	}

	for _, col := range in.Table.Columns {
		if len(col.Id) == 0 {
			col.Id = ksuid.New().String()
		}
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

	return model.GetTableModel().Get(ctx, in.Id)
}

// ListTables implements Interface.
func (s unitableServer) ListTables(ctx context.Context, in *pb.ListTablesRequest) (*pb.ListTablesResponse, error) {
	if len(in.Database) == 0 {
		return nil, core.NewInvalidArgumentError("not set the database")
	}

	tables, err := model.GetTableModel().Query(ctx)
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

	if _, err := s.Synchro.InsertRow(ctx, in.Table, in.Row); err != nil {
		return nil, core.NewInternalError("failed to insert the row in %s, (%v)", in.Table, in.Row.ToMapInterface())
	}

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

	if _, err := s.Synchro.UpdateRow(ctx, in.Table, in.Row); err != nil {
		return nil, core.NewInternalError("failed to update the row in %s, (%v)", in.Table, in.Row.ToMapInterface())
	}

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

	if row, err := s.Synchro.GetRow(ctx, in.Table, in.Id); err != nil {
		return nil, core.NewInternalError("failed to get the row %s in %s", in.Id, in.Table)
	} else {
		return row, nil
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

	if _, err := s.Synchro.DeleteRows(ctx, in.Table, in.Id); err != nil {
		return nil, core.NewInternalError("failed to delete the row %s in %s", in.Id, in.Table)
	} else {
		return &core.Null{}, nil
	}
}

// ListRow implements Interface.
func (s unitableServer) ListRow(ctx context.Context, in *pb.ListRowRequest) (*pb.ListRowResponse, error) {
	resp := &pb.ListRowResponse{
		// Objects:
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}

// ExportRow implements Interface.
func (s unitableServer) ExportRow(ctx context.Context, in *pb.ExportRowRequest) (*pb.ExportRowResponse, error) {
	resp := &pb.ExportRowResponse{
		// Objects:
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}

// BatchCreateRows implements Interface.
func (s unitableServer) BatchCreateRows(ctx context.Context, in *pb.BatchCreateRowsRequest) (*core.Null, error) {
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
		id := row.GetString("id")
		if len(id) == 0 {
			id = ksuid.New().String()
			row.SetString("id", id)
		}
	}

	if _, err := s.Synchro.InsertRows(ctx, in.Table, in.Rows...); err != nil {
		return nil, core.NewInternalError("failed to batch create the rows in %s, err: %s", in.Table, err.Error())
	}

	return &core.Null{}, nil
}

// BatchUpdateRows implements Interface.
func (s unitableServer) BatchUpdateRows(ctx context.Context, in *pb.BatchUpdateRowsRequest) (*core.Null, error) {
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
		id := row.GetString("id")
		if len(id) == 0 {
			return nil, core.NewInvalidArgumentError("the No. %d (begin with 1) row have not set the id in batch", i+1)
		}
	}

	if _, err := s.Synchro.UpdateRows(ctx, in.Table, in.Rows...); err != nil {
		return nil, core.NewInternalError("failed to batch update the rows in %s, err: %s", in.Table, err.Error())
	}

	return &core.Null{}, nil
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

	if _, err := s.Synchro.DeleteRows(ctx, in.Table, in.Ids...); err != nil {
		return nil, core.NewInternalError("failed to batch delete the row in %s, err: %s", in.Table, err.Error())
	}

	return &core.Null{}, nil
}
