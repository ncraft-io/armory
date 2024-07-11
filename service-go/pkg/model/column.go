package model

import (
	"context"
	"sync"

	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/gorm"
)

var clmn *Column
var clmnOnce sync.Once

type Column struct {
	DB *db.DB
}

func GetColumnModel() *Column {
	clmnOnce.Do(func() {
		clmn = NewColumn()
	})

	return clmn
}

func NewColumn() *Column {
	t := &Column{DB: GetDB()}
	if !t.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&unitable.Column{}) {
		if err := d.AutoMigrate(&unitable.Column{}); err != nil {
			logs.ErrLog("Init Column model err: ", err)
			panic(err)
		}
	}
	return t
}

func (a *Column) Create(ctx context.Context, columns ...*unitable.Column) (int64, error) {
	columnsLen := len(columns)
	var executionResult *gorm.DB

	if columnsLen == 0 {
		return 0, nil
	} else if columnsLen == 1 {
		//executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(columns[0])
		executionResult = a.DB.WithContext(ctx).Create(columns[0])
	} else {
		executionResult = a.DB.WithContext(ctx).CreateInBatches(columns, len(columns))
	}

	return executionResult.RowsAffected, executionResult.Error
}

func (a *Column) Get(ctx context.Context, uid string) (*unitable.Column, error) {
	column := &unitable.Column{}
	return column, a.DB.WithContext(ctx).First(column, "id = ?", uid).Error
}

func (a *Column) BatchGet(ctx context.Context, ids []string) ([]*unitable.Column, error) {
	var columns []*unitable.Column
	return columns, a.DB.WithContext(ctx).Find(&columns, ids).Error
}

func (a *Column) Query(ctx context.Context) ([]*unitable.Column, error) {
	var columns []*unitable.Column

	tx := a.DB.DB.WithContext(ctx)
	//if len(projectId) > 0 {
	//	tx = tx.Where("project_id = ?", projectId)
	//}
	//if len(sourcePackageId) > 0 {
	//	tx = tx.Where("source_package_id = ?", sourcePackageId)
	//}
	//if len(buildId) > 0 {
	//	tx = tx.Where("build_id = ?", buildId)
	//}
	//if len(operationName) > 0 {
	//	tx = tx.Where("operation_name = ?", operationName)
	//}
	//if len(buildOperationName) > 0 {
	//	tx = tx.Where("build_operation_name = ?", buildOperationName)
	//}
	//if len(name) > 0 {
	//	tx = tx.Where("name = ?", name)
	//}

	return columns, tx.Find(&columns).Error
}

func (a *Column) Delete(ctx context.Context, uid string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Where("id = ?", uid).Delete(&unitable.Column{})
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Column) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Delete(&unitable.Column{}, ids)
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Column) Update(ctx context.Context, column *unitable.Column) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Updates(column)
	return executionResult.RowsAffected, executionResult.Error
}
