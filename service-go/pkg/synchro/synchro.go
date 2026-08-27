package synchro

import (
	"context"
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/model"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"sync"
)

type MetaTable struct {
	Table      *unitable.Table
	Struct     *DynamicStruct
	FieldsInfo query.Fields
}

type Synchro struct {
	Tables sync.Map //map[string]*MetaTable
}

func TableId(database string, table string) string {
	return strings.Join([]string{database, table}, ".")
}

func New() *Synchro {
	return &Synchro{}
}

func (s *Synchro) GetMetaTable(tableId string, table *unitable.Table) *MetaTable {
	if table == nil { // if table not nil, will using user table, and not cache
		if value, ok := s.Tables.Load(tableId); ok {
			if mt, ok := value.(*MetaTable); ok {
				return mt
			}
		}
	}

	tb := table
	if tb == nil {
		t, err := model.GetTableModel().Get(context.Background(), tableId)
		if err != nil {
			logs.ErrLogw("failed to get the table", "name", table, "error", err)
			return nil
		}
		if len(t.GetColumns()) == 0 {
			logs.ErrLogw("failed to get any columns in the table", "name", table)
			return nil
		}

		tb = t
	}

	meta := &MetaTable{
		Table:      tb,
		Struct:     NewDynamicStruct(tb),
		FieldsInfo: make(query.Fields),
	}
	for _, col := range tb.Columns {
		meta.FieldsInfo[col.Name] = col.ToFieldInfo()
	}

	logs.Infow("get the meta table ok", "database", tb.Database, "table", tb.Name, "columns", len(tb.Columns))

	if table == nil {
		s.Tables.Store(tableId, meta)
	}
	return meta
}

const createSql = `CREATE TABLE %s (id varchar(255) NOT NULL PRIMARY KEY);`

func (s *Synchro) CreateTable(ctx context.Context, table *unitable.Table) error {
	if !GetDataDB(table.Database).Migrator().HasTable(table.Id) {
		tx := GetDataDB(table.Database).WithContext(ctx).Exec(fmt.Sprintf(createSql, table.Id))
		return tx.Error
	}
	return nil
}

func (s *Synchro) MigrateTable(ctx context.Context, table *unitable.Table, renamedCols map[string]string, dropCols []string) error {
	var meta *MetaTable
	if len(renamedCols) > 0 {
		meta = s.GetMetaTable(table.Id, nil)
		if meta == nil {
			return core.NewNotFoundError("the original table %s is not exist, can't to rename fields %v", table.Name, renamedCols)
		}

		obj := meta.Struct.New()
		tx := GetDataDB(table.Database).WithContext(ctx).Table(table.Name)
		for k, v := range renamedCols {
			if err := tx.Migrator().RenameColumn(obj, k, v); err != nil {
				return err
			}
		}
	}

	{
		s.Tables.Delete(table.Id)
	}

	meta = s.GetMetaTable(table.Id, table)
	if meta == nil {
		return core.NewNotFoundError("the table %s is not exist", table.Name)
	}

	return GetDataDB(table.Database).WithContext(ctx).Table(table.Name).AutoMigrate(meta.Struct.New())
}

func (s *Synchro) DropTable(ctx context.Context, table *unitable.Table) error {
	meta := s.GetMetaTable(table.Id, nil)
	if meta == nil {
		return core.NewNotFoundError("the table %s is not exist", table.Id)
	}

	{
		s.Tables.Delete(table.Id)
	}

	tx := GetDataDB(table.Database).WithContext(ctx).Exec("DROP TABLE " + table.Id)
	return tx.Error
}

func (s *Synchro) GetRow(ctx context.Context, table string, id string) (*core.Object, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	row := meta.Struct.New()
	resul := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).First(row, "id = ?", id)
	if resul.Error != nil {
		return nil, core.NewNotFoundError("%s is not exist in %s, %s", id, table, resul.Error.Error())
	}

	obj, err := ParseObject(row)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *Synchro) BatchGetRow(ctx context.Context, table string, ids ...string) ([]*core.Object, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return nil, core.NewNotFoundError("the table %s is not exist", table)
	}
	if len(ids) == 0 {
		return nil, nil
	}

	rows := meta.Struct.NewSliceOf()
	resul := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Where(ids).Find(rows)
	if resul.Error != nil {
		return nil, core.NewNotFoundError("%s is not exist in %s, %s", ids, table, resul.Error.Error())
	}

	objs, err := ParseObjects(rows)
	if err != nil {
		return nil, err
	}

	return objs, nil
}

func (s *Synchro) QueryBy(ctx context.Context, database string, tableName string, query *unitable.DbQuery, arguments []interface{}) ([]*core.Object, error) {
	table := &unitable.Table{
		Database:  database,
		Name:      tableName,
		JsonStyle: query.JsonStyle,
		Columns:   query.Columns,
	}
	meta := &MetaTable{
		Table:  table,
		Struct: NewDynamicStruct(table),
	}

	tx := GetDataDB(database).WithContext(ctx).Table(table.Name).Raw(query.Sql, arguments...)
	rows, err := tx.Rows()
	defer rows.Close()
	if err != nil {
		return nil, core.NewNotFoundError("failed to query the rows in %s, %s", table, err.Error())
	}

	var objs []*core.Object
	for rows.Next() {
		row := meta.Struct.New()
		if err = GetDataDB(database).ScanRows(rows, row); err != nil {
			return nil, core.NewInternalError("failed to scan the row in %s, %s", table, err.Error())
		}

		obj, err := ParseObject(row)
		if err != nil {
			return nil, core.NewInternalError("failed to parse the row in %s, %s", table, err.Error())
		}
		objs = append(objs, obj)
	}

	return objs, nil
}

func (s *Synchro) QueryRows(ctx context.Context, table string, query *query.Query) ([]*core.Object, int, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil || meta.Table == nil {
		logs.Warnw("the table is not exist", "table", table)
		return nil, 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name)
	if query != nil {
		tx = query.Apply(tx, meta.FieldsInfo)
	} else {
		tx = tx.Select("*")
	}

	rows, err := tx.Rows()
	defer rows.Close()
	if err != nil {
		return nil, 0, core.NewNotFoundError("failed to query the rows in %s, %s", table, err.Error())
	}

	st := NewDynamicStructWith(query, meta)

	var objs []*core.Object
	for rows.Next() {
		row := st.New()
		if err = GetDataDB(meta.Table.Database).ScanRows(rows, row); err != nil {
			return nil, 0, err
		}

		obj, err := ParseObject(row)
		if err != nil {
			return nil, 0, err
		}
		objs = append(objs, obj)
	}

	totalCnt := len(objs)
	if query.PageSize > 0 { // pagination
		// get the total count
		tx = GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name)
		if query != nil {
			tx = query.TotalCount(tx, meta.FieldsInfo)
		} else {
			tx = tx.Select("COUNT(*)")
		}
		row := tx.Row()

		if row != nil {
			err = row.Scan(&totalCnt)
			if err != nil {
				return nil, 0, core.NewNotFoundError("failed to query the total count in %s, %s", table, err.Error())
			}
		}
	}

	return objs, totalCnt, nil
}

func (s *Synchro) InsertRow(ctx context.Context, table string, row *core.Object) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}
	logs.Debugf("the meta table is %v", meta.Table)

	data, err := meta.Struct.NewOf(row)
	if err != nil {
		return 0, err
	}
	logs.Debugf("the create row is %v", data)

	result := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Clauses(clause.OnConflict{UpdateAll: true}).Create(data)
	return result.RowsAffected, result.Error
}

func (s *Synchro) InsertRows(ctx context.Context, table string, rows ...*core.Object) (int64, []int, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	irs := make([]int, len(rows))

	// Continuous session mode
	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})
	err := tx.Transaction(func(tx *gorm.DB) error {
		for i, row := range rows {
			data, err := meta.Struct.NewOf(row)
			if err != nil {
				return err
			}
			result := tx.Table(meta.Table.Name).Clauses(clause.OnConflict{UpdateAll: true}).Create(data)
			if result.Error != nil {
				// return any error will roll back
				logs.Warnw("failed to create the row, will rollback", "row", row, "error", err)
				return err
			}
			if result.RowsAffected > 0 {
				irs[i] = 1
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return 0, nil, err
	}
	return int64(len(rows)), irs, nil
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

	updateRow := FilterOutId(row.ToMapInterface(), meta.Table)
	result := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Model(data).Updates(updateRow)
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
	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})
	err := tx.Transaction(func(tx *gorm.DB) error {
		for i, data := range datas {
			updateRow := FilterOutId(rows[i].ToMapInterface(), meta.Table)
			if err := tx.Table(meta.Table.Name).Model(data).Updates(updateRow).Error; err != nil {
				// return any error will roll back
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	return int64(len(datas)), err
}

func (s *Synchro) UpdateInsertRows(ctx context.Context, table string, rows ...*core.Object) (int64, []int, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	var data []interface{}
	var dataIndex []int
	var insertData []interface{}
	var insertDataIndex []int
	for i, row := range rows {
		id := row.GetString("id")
		if len(id) == 0 {
			row.SetString("id", ksuid.New().String())
			r, err := meta.Struct.NewOf(row)
			if err != nil {
				return 0, nil, err
			}
			insertData = append(insertData, r)
			insertDataIndex = append(insertDataIndex, i)
		} else {
			r, err := meta.Struct.NewOf(row)
			if err != nil {
				return 0, nil, err
			}
			data = append(data, r)
			dataIndex = append(dataIndex, i)
		}
	}

	irs := make([]int, len(rows))

	// Continuous session mode
	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Session(&db.Session{SkipDefaultTransaction: true})
	err := tx.Transaction(func(tx *gorm.DB) error {
		for i, d := range insertData {
			result := tx.Table(meta.Table.Name).Create(d)
			if result.Error != nil {
				logs.ErrLogw("failed to create the row, will rollback", "index", insertDataIndex[i], "error", result.Error)
				return result.Error
			}

			if result.RowsAffected > 0 {
				irs[insertDataIndex[i]] = 1
			}
		}

		for i, d := range data {
			result := tx.Table(meta.Table.Name).Updates(d)
			if result.Error != nil {
				logs.ErrLogw("failed to update the row, will rollback", "index", dataIndex[i], "error", result.Error)
				return result.Error
			}

			if result.RowsAffected > 0 {
				irs[dataIndex[i]] = 1
			}
		}

		// return nil will commit the whole transaction
		return nil
	})

	return int64(len(data) + len(insertData)), irs, err
}

func (s *Synchro) DeleteRows(ctx context.Context, table string, ids ...string) (int64, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return 0, core.NewNotFoundError("the table %s is not exist", table)
	}

	result := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name).Delete(meta.Struct.New(), ids)
	return result.RowsAffected, result.Error
}

func (s *Synchro) CalcStats(ctx context.Context, table string, query *query.Query) ([]*core.Object, error) {
	meta := s.GetMetaTable(table, nil)
	if meta == nil {
		return nil, core.NewNotFoundError("the table %s is not exist", table)
	}

	tx := GetDataDB(meta.Table.Database).WithContext(ctx).Table(meta.Table.Name)
	if query == nil {
		return nil, core.NewInvalidArgumentError("the query is null from table %s", table)
	}

	tx = query.Apply(tx, meta.FieldsInfo)

	rows, err := tx.Rows()
	defer rows.Close()
	if err != nil {
		return nil, core.NewNotFoundError("failed to query the rows in %s, %s", table, err.Error())
	}

	st := NewDynamicStructWith(query, meta)

	var objs []*core.Object
	for rows.Next() {
		row := st.New()
		if err = GetDataDB(meta.Table.Database).ScanRows(rows, row); err != nil {
			return nil, err
		}

		obj, err := ParseObject(row)
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}
