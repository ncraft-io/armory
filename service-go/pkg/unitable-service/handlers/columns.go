package handlers

import (
	"context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"google.golang.org/protobuf/proto"
)

func findColumn(table *unitable.Table, id string) (*unitable.Column, error) {
	if id == "" {
		return nil, core.NewInvalidArgumentError("not set the column id")
	}
	for _, col := range table.Columns {
		if col.Id == id {
			return col, nil
		}
	}
	return nil, core.NewNotFoundError("column %s not found in table %s", id, table.Id)
}

func (s unitableServer) CreateColumn(ctx context.Context, in *pb.CreateColumnRequest) (*unitable.Column, error) {
	if in == nil || in.Column == nil {
		return nil, core.NewInvalidArgumentError("not set the column")
	}
	columns, err := s.createColumns(ctx, in.Database, in.Table, []*unitable.Column{in.Column})
	if err != nil {
		return nil, err
	}
	return columns[0], nil
}

func (s unitableServer) createColumns(ctx context.Context, database, name string, columns []*unitable.Column) ([]*unitable.Column, error) {
	if len(columns) == 0 {
		return nil, core.NewInvalidArgumentError("not set the columns")
	}
	s.instance.schemaMu.Lock()
	defer s.instance.schemaMu.Unlock()
	table, err := s.getTable(ctx, database, name)
	if err != nil {
		return nil, err
	}
	created := make([]*unitable.Column, 0, len(columns))
	for _, col := range columns {
		if col == nil {
			return nil, core.NewInvalidArgumentError("nil column")
		}
		created = append(created, proto.Clone(col).(*unitable.Column))
	}
	table.Columns = append(table.Columns, created...)
	if err := validateColumns(table); err != nil {
		return nil, err
	}
	prepareColumns(table)
	if err := s.checkColumnIDs(ctx, table); err != nil {
		return nil, err
	}
	if err := s.Synchro().MigrateTable(ctx, table, nil, nil); err != nil {
		return nil, err
	}
	table.UpdateTime = core.Now()
	if err := s.saveTable(ctx, table); err != nil {
		return nil, err
	}
	return created, nil
}

func (s unitableServer) UpdateColumn(ctx context.Context, in *pb.UpdateColumnRequest) (*core.Null, error) {
	if in == nil || in.Column == nil {
		return nil, core.NewInvalidArgumentError("not set the column")
	}
	col := proto.Clone(in.Column).(*unitable.Column)
	if in.Id != "" {
		if col.Id != "" && col.Id != in.Id {
			return nil, core.NewInvalidArgumentError("column id differs from request path")
		}
		col.Id = in.Id
	}
	return s.updateColumns(ctx, in.Database, in.Table, []*unitable.Column{col})
}

func (s unitableServer) updateColumns(ctx context.Context, database, name string, columns []*unitable.Column) (*core.Null, error) {
	if len(columns) == 0 {
		return nil, core.NewInvalidArgumentError("not set the columns")
	}
	s.instance.schemaMu.Lock()
	defer s.instance.schemaMu.Unlock()
	old, err := s.getTable(ctx, database, name)
	if err != nil {
		return nil, err
	}
	table := proto.Clone(old).(*unitable.Table)
	table.Columns = nil
	for _, col := range columns {
		if col == nil {
			return nil, core.NewInvalidArgumentError("nil column")
		}
		if _, err := findColumn(old, col.Id); err != nil {
			return nil, err
		}
		table.Columns = append(table.Columns, proto.Clone(col).(*unitable.Column))
	}
	if _, err := s.updateTable(ctx, table, old, false); err != nil {
		return nil, err
	}
	return &core.Null{}, nil
}

func (s unitableServer) GetColumn(ctx context.Context, in *pb.GetColumnRequest) (*unitable.Column, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	table, err := s.getTable(ctx, in.Database, in.Table)
	if err != nil {
		return nil, err
	}
	return findColumn(table, in.Id)
}

func (s unitableServer) DeleteColumn(ctx context.Context, in *pb.DeleteColumnRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	return s.deleteColumns(ctx, in.Database, in.Table, []string{in.Id})
}

func (s unitableServer) deleteColumns(ctx context.Context, database, name string, ids []string) (*core.Null, error) {
	if len(ids) == 0 {
		return nil, core.NewInvalidArgumentError("not set the column ids")
	}
	s.instance.schemaMu.Lock()
	defer s.instance.schemaMu.Unlock()
	table, err := s.getTable(ctx, database, name)
	if err != nil {
		return nil, err
	}
	deleting := map[string]bool{}
	var drop []string
	for _, id := range ids {
		col, err := findColumn(table, id)
		if err != nil {
			return nil, err
		}
		if col.Name == "id" {
			return nil, core.NewInvalidArgumentError("cannot delete the row id column")
		}
		if !deleting[id] {
			drop = append(drop, col.Name)
		}
		deleting[id] = true
	}
	var keep []*unitable.Column
	for _, col := range table.Columns {
		if !deleting[col.Id] {
			keep = append(keep, col)
		}
	}
	if len(keep) == 0 {
		return nil, core.NewInvalidArgumentError("cannot delete all columns")
	}
	table.Columns = keep
	if err := s.Synchro().MigrateTable(ctx, table, nil, drop); err != nil {
		return nil, err
	}
	table.UpdateTime = core.Now()
	if err := s.saveTable(ctx, table); err != nil {
		return nil, err
	}
	return &core.Null{}, nil
}

func (s unitableServer) ListColumns(ctx context.Context, in *pb.ListColumnsRequest) (*pb.ListColumnsResponse, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	table, err := s.getTable(ctx, in.Database, in.Table)
	if err != nil {
		return nil, err
	}
	qry, err := scopedColumnQuery(in, table.Id)
	if err != nil {
		return nil, err
	}
	columns, err := model.GetColumnModel().List(ctx, qry)
	if err != nil {
		return nil, err
	}
	allQuery, _ := scopedColumnQuery(in, table.Id)
	allQuery.PageSize, allQuery.PageToken, allQuery.Skip = 0, "", 0
	all, err := model.GetColumnModel().List(ctx, allQuery)
	if err != nil {
		return nil, err
	}
	return &pb.ListColumnsResponse{Columns: columns, TotalCount: int32(len(all)), NextPageToken: nextPageToken(qry, len(all))}, nil
}

func (s unitableServer) BatchCreateColumns(ctx context.Context, in *pb.BatchCreateColumnsRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	if _, err := s.createColumns(ctx, in.Database, in.Table, in.Columns); err != nil {
		return nil, err
	}
	return &core.Null{}, nil
}
func (s unitableServer) BatchUpdateColumn(ctx context.Context, in *pb.BatchUpdateColumnRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	return s.updateColumns(ctx, in.Database, in.Table, in.Columns)
}
func (s unitableServer) BatchDeleteColumn(ctx context.Context, in *pb.BatchDeleteColumnRequest) (*core.Null, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	return s.deleteColumns(ctx, in.Database, in.Table, in.Ids)
}
