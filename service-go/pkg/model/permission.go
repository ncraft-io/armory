package model

import (
	"context"
	"gorm.io/gorm/clause"
	"sync"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/ncraft-io/armory/go/pkg/armory/auth"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/gorm"
)

var permission *Permission
var permissionOnce sync.Once

type Permission struct {
	DB *db.DB
}

func GetPermissionModel() *Permission {
	permissionOnce.Do(func() {
		permission = NewPermission()
	})

	return permission
}

func NewPermission() *Permission {
	t := &Permission{DB: GetDB()}
	if !t.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&auth.Permission{}) {
		if err := d.AutoMigrate(&auth.Permission{}); err != nil {
			logs.ErrLog("Init Permission model err: ", err)
			panic(err)
		}
	}
	return t
}

func (a *Permission) Create(ctx context.Context, permissions ...*auth.Permission) (int64, error) {
	permissionsLen := len(permissions)
	var executionResult *gorm.DB

	if permissionsLen == 0 {
		return 0, nil
	} else if permissionsLen == 1 {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(permissions[0])
	} else {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(permissions, len(permissions))
	}

	return executionResult.RowsAffected, executionResult.Error
}

func (a *Permission) Get(ctx context.Context, uid string) (*auth.Permission, error) {
	permission := &auth.Permission{}
	return permission, a.DB.WithContext(ctx).First(permission, "id = ?", uid).Error
}

func (a *Permission) BatchGet(ctx context.Context, ids []string) ([]*auth.Permission, error) {
	var permissions []*auth.Permission
	return permissions, a.DB.WithContext(ctx).Find(&permissions, ids).Error
}

func (a *Permission) Query(ctx context.Context) ([]*auth.Permission, error) {
	var permissions []*auth.Permission

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

	return permissions, tx.Find(&permissions).Error
}

func (a *Permission) Delete(ctx context.Context, uid string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Where("id = ?", uid).Delete(&auth.Permission{})
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Permission) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Delete(&auth.Permission{}, ids)
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Permission) Update(ctx context.Context, permission *auth.Permission) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Updates(permission)
	return executionResult.RowsAffected, executionResult.Error
}
