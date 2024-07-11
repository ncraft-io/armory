package synchro

import (
	"context"
	"fmt"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/gorm"
)

type MetaTable struct {
	Table  *unitable.Table
	Struct *DynamicStruct
}

type Synchro struct {
	Tables map[string]*MetaTable
}

func New() *Synchro {
	return &Synchro{
		Tables: make(map[string]*MetaTable),
	}
}

func (s *Synchro) GetMetaTable(table string) *MetaTable {
	if mt, ok := s.Tables[table]; !ok {
		t, err := model.GetTableModel().Get(context.Background(), table)
		if err != nil {
			logs.ErrLogw("failed to get the table", "name", table, "error", err)
			return nil
		}

		meta := &MetaTable{
			Table:  t,
			Struct: NewDynamicStruct(t),
		}

		s.Tables[table] = meta
		return meta
	} else {
		return mt
	}
}

const createSql = `CREATE TABLE %s (id varchar(255) NOT NULL PRIMARY KEY);`

func (s *Synchro) CreateTable(ctx context.Context, table *unitable.Table) error {
	if !GetDataDB().Migrator().HasTable(table.Id) {
		tx := GetDataDB().WithContext(ctx).Exec(fmt.Sprintf(createSql, table.Id))
		return tx.Error
	}
	return nil
}

func (s *Synchro) MigrateTable(ctx context.Context, table *unitable.Table) error {
	meta := s.GetMetaTable(table.Id)
	if meta == nil {
		return core.NewNotFoundError("the table %s is not exist", table.Name)
	}

	delete(s.Tables, table.Id)
	return GetDataDB().WithContext(ctx).Table(table.Id).AutoMigrate(meta.Struct.New())
}

func (s *Synchro) DropTable(ctx context.Context, table *unitable.Table) error {
	meta := s.GetMetaTable(table.Id)
	if meta == nil {
		return core.NewNotFoundError("the table %s is not exist", table.Id)
	}

	delete(s.Tables, table.Id)
	tx := GetDataDB().WithContext(ctx).Exec("DROP TABLE " + table.Id)
	return tx.Error
}

func (s *Synchro) GetRow(ctx context.Context, table string, id string) (*core.Object, error) {
	meta := s.GetMetaTable(table)
	if meta == nil {
		return nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	row := meta.Struct.New()
	resul := GetDataDB().WithContext(ctx).Table(table).First(row, "id = ?", id)
	if resul.Error != nil {
		return nil, core.NewNotFoundError("%s is not exist in %s, %s", id, table, resul.Error.Error())
	}

	obj, err := ParseObject(row)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *Synchro) QueryRows(ctx context.Context, table string) ([]*core.Object, error) {
	return nil, nil
}

func (s *Synchro) InsertRow(ctx context.Context, table string, row *core.Object) (int64, error) {
	meta := s.GetMetaTable(table)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	data, err := meta.Struct.NewOf(row)
	if err != nil {
		return 0, err
	}

	result := GetDataDB().WithContext(ctx).Table(table).Create(data)
	return result.RowsAffected, result.Error
}

func (s *Synchro) InsertRows(ctx context.Context, table string, rows ...*core.Object) (int64, error) {
	meta := s.GetMetaTable(table)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	var data []interface{}
	for _, row := range rows {
		r, err := meta.Struct.NewOf(row)
		if err != nil {
			return 0, err
		}
		data = append(data, r)
	}

	result := GetDataDB().WithContext(ctx).Table(table).CreateInBatches(data, len(data))
	return result.RowsAffected, result.Error
}

func (s *Synchro) UpdateRow(ctx context.Context, table string, row *core.Object) (int64, error) {
	meta := s.GetMetaTable(table)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	data, err := meta.Struct.NewOf(row)
	if err != nil {
		return 0, err
	}

	result := GetDataDB().WithContext(ctx).Table(table).Updates(data)
	return result.RowsAffected, result.Error
}

func (s *Synchro) UpdateRows(ctx context.Context, table string, rows ...*core.Object) (int64, error) {
	meta := s.GetMetaTable(table)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	var data []interface{}
	for _, row := range rows {
		r, err := meta.Struct.NewOf(row)
		if err != nil {
			return 0, err
		}
		data = append(data, r)
	}

	// Continuous session mode
	tx := GetDataDB().WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})

	err := tx.Transaction(func(tx *gorm.DB) error {
		for _, d := range data {
			if err := tx.Table(table).Updates(d).Error; err != nil {
				// return any error will roll back
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	return int64(len(data)), err
}

func (s *Synchro) DeleteRows(ctx context.Context, table string, ids ...string) (int64, error) {
	meta := s.GetMetaTable(table)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	result := GetDataDB().WithContext(ctx).Table(table).Delete(meta.Struct.New(), ids)
	return result.RowsAffected, result.Error
}
