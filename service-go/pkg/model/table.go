package model

import (
	"context"
	"sync"

	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/gorm"
)

var tbl *Table
var tblOnce sync.Once

type Table struct {
	DB *db.DB
}

func GetTableModel() *Table {
	tblOnce.Do(func() {
		tbl = NewTable()
	})

	return tbl
}

func NewTable() *Table {
	t := &Table{DB: GetDB()}
	if !t.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&unitable.Table{}) {
		if err := d.AutoMigrate(&unitable.Table{}); err != nil {
			logs.ErrLog("Init Table model err: ", err)
			panic(err)
		}

		if err := d.AutoMigrate(&unitable.Column{}); err != nil {
			logs.ErrLog("Init Column model err: ", err)
			panic(err)
		}
	}
	return t
}

func (a *Table) Create(ctx context.Context, tables ...*unitable.Table) (int64, error) {
	tablesLen := len(tables)
	var executionResult *gorm.DB

	if tablesLen == 0 {
		return 0, nil
	} else if tablesLen == 1 {
		//executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(tables[0])
		executionResult = a.DB.WithContext(ctx).Create(tables[0])
	} else {
		executionResult = a.DB.WithContext(ctx).CreateInBatches(tables, len(tables))
	}

	return executionResult.RowsAffected, executionResult.Error
}

func (a *Table) Get(ctx context.Context, uid string) (*unitable.Table, error) {
	table := &unitable.Table{}
	return table, a.DB.WithContext(ctx).Preload("Columns").First(table, "id = ?", uid).Error
}

func (a *Table) BatchGet(ctx context.Context, ids []string) ([]*unitable.Table, error) {
	var tables []*unitable.Table
	return tables, a.DB.WithContext(ctx).Find(&tables, ids).Error
}

func (a *Table) Query(ctx context.Context) ([]*unitable.Table, error) {
	var tables []*unitable.Table

	tx := a.DB.DB.WithContext(ctx)

	return tables, tx.Find(&tables).Error
}

func (a *Table) Delete(ctx context.Context, uid string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Where("id = ?", uid).Delete(&unitable.Table{})
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Table) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Delete(&unitable.Table{}, ids)
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Table) Update(ctx context.Context, table *unitable.Table) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Updates(table)
	return executionResult.RowsAffected, executionResult.Error
}
