package synchro

import (
	"context"
	"fmt"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/segmentio/ksuid"
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

func (s *Synchro) GetMetaTable(tableName string, table *unitable.Table) *MetaTable {
	if mt, ok := s.Tables[tableName]; !ok {
		if table == nil {
			t, err := model.GetTableModel().Get(context.Background(), tableName)
			if err != nil {
				logs.ErrLogw("failed to get the table", "name", table, "error", err)
				return nil
			}
			table = t
		}

		meta := &MetaTable{
			Table:  table,
			Struct: NewDynamicStruct(table),
		}

		// s.Tables[tableName] = meta
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

func (s *Synchro) MigrateTable(ctx context.Context, table *unitable.Table, renamedCols map[string]string) error {
	meta := s.GetMetaTable(table.Id, table)
	if meta == nil {
		return core.NewNotFoundError("the table %s is not exist", table.Name)
	}

	delete(s.Tables, table.Id)

	if renamedCols != nil {
		obj := meta.Struct.New()
		tx := GetDataDB().WithContext(ctx).Table(table.Id)
		for k, v := range renamedCols {
			if err := tx.Migrator().RenameColumn(obj, k, v); err != nil {
				return err
			}
		}
	}

	return GetDataDB().WithContext(ctx).Table(table.Id).AutoMigrate(meta.Struct.New())
}

func (s *Synchro) DropTable(ctx context.Context, table *unitable.Table) error {
	meta := s.GetMetaTable(table.Id, nil)
	if meta == nil {
		return core.NewNotFoundError("the table %s is not exist", table.Id)
	}

	delete(s.Tables, table.Id)
	tx := GetDataDB().WithContext(ctx).Exec("DROP TABLE " + table.Id)
	return tx.Error
}

func (s *Synchro) GetRow(ctx context.Context, table string, id string) (*core.Object, error) {
	meta := s.GetMetaTable(table, nil)
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

func (s *Synchro) QueryRows(ctx context.Context, table string, query *db.Query) ([]*core.Object, int, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return nil, 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	tx := GetDataDB().WithContext(ctx).Table(table)
	if query != nil {
		tx = query.Apply(tx)
	} else {
		tx = tx.Select("*")
	}

	rows, err := tx.Rows()
	defer rows.Close()
	if err != nil {
		return nil, 0, core.NewNotFoundError("failed to query the rows in %s, %s", table, err.Error())
	}

	var objs []*core.Object
	for rows.Next() {
		row := meta.Struct.New()
		if err = GetDataDB().ScanRows(rows, row); err != nil {
			return nil, 0, err
		}

		obj, err := ParseObject(row)
		if err != nil {
			return nil, 0, err
		}
		objs = append(objs, obj)
	}

	// get the total count
	tx = GetDataDB().WithContext(ctx).Table(table)
	if query != nil {
		tx = query.ApplyTotalCount(tx)
	} else {
		tx = tx.Select("COUNT(*)")
	}
	row := tx.Row()
	totalCnt := 0
	if row != nil {
		err = row.Scan(&totalCnt)
		if err != nil {
			return nil, 0, core.NewNotFoundError("failed to query the total count in %s, %s", table, err.Error())
		}
	}

	return objs, totalCnt, nil
}

func (s *Synchro) InsertRow(ctx context.Context, table string, row *core.Object) (int64, error) {
	meta := s.GetMetaTable(table, nil)
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
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	// Continuous session mode
	tx := GetDataDB().WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})
	err := tx.Transaction(func(tx *gorm.DB) error {
		for _, row := range rows {
			data, err := meta.Struct.NewOf(row)
			if err != nil {
				return err
			}
			if err := tx.Table(table).Create(data).Error; err != nil {
				// return any error will roll back
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return 0, err
	}
	return int64(len(rows)), nil

	//result := GetDataDB().WithContext(ctx).Table(table).CreateInBatches(data, len(data))
	//return result.RowsAffected, result.Error
}

func (s *Synchro) UpdateRow(ctx context.Context, table string, row *core.Object) (int64, error) {
	meta := s.GetMetaTable(table, nil)
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
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	var datas []interface{}
	for _, row := range rows {
		r, err := meta.Struct.NewOf(row)
		if err != nil {
			return 0, err
		}
		datas = append(datas, r)
	}

	// Continuous session mode
	tx := GetDataDB().WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})
	err := tx.Transaction(func(tx *gorm.DB) error {
		for _, date := range datas {
			if err := tx.Table(table).Updates(date).Error; err != nil {
				// return any error will roll back
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	return int64(len(datas)), err
}

func (s *Synchro) UpdateInsertRows(ctx context.Context, table string, rows ...*core.Object) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	var data []interface{}
	var insertData []interface{}
	for _, row := range rows {
		id := row.GetString("id")
		if len(id) == 0 {
			row.SetString("id", ksuid.New().String())
			r, err := meta.Struct.NewOf(row)
			if err != nil {
				return 0, err
			}
			insertData = append(insertData, r)
		} else {
			r, err := meta.Struct.NewOf(row)
			if err != nil {
				return 0, err
			}
			data = append(data, r)
		}
	}

	// Continuous session mode
	tx := GetDataDB().WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})

	err := tx.Transaction(func(tx *gorm.DB) error {
		for _, d := range insertData {
			if err := tx.Table(table).Create(d).Error; err != nil {
				return err
			}
		}

		for _, d := range data {
			if err := tx.Table(table).Updates(d).Error; err != nil {
				// return any error will roll back
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	return int64(len(data) + len(insertData)), err
}

func (s *Synchro) DeleteRows(ctx context.Context, table string, ids ...string) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	result := GetDataDB().WithContext(ctx).Table(table).Delete(meta.Struct.New(), ids)
	return result.RowsAffected, result.Error
}
