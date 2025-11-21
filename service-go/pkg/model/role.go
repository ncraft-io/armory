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

var role *Role
var roleOnce sync.Once

type Role struct {
	DB *db.DB
}

func GetRoleModel() *Role {
	roleOnce.Do(func() {
		role = NewRole()
	})

	return role
}

func NewRole() *Role {
	t := &Role{DB: GetDB()}
	if !t.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&auth.Role{}) {
		if err := d.AutoMigrate(&auth.Role{}); err != nil {
			logs.ErrLog("Init Role model err: ", err)
			panic(err)
		}
	}
	return t
}

func (a *Role) Create(ctx context.Context, roles ...*auth.Role) (int64, error) {
	rolesLen := len(roles)
	var executionResult *gorm.DB

	if rolesLen == 0 {
		return 0, nil
	} else if rolesLen == 1 {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(roles[0])
	} else {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(roles, len(roles))
	}

	return executionResult.RowsAffected, executionResult.Error
}

func (a *Role) Get(ctx context.Context, uid string) (*auth.Role, error) {
	role := &auth.Role{}
	return role, a.DB.WithContext(ctx).First(role, "id = ?", uid).Error
}

func (a *Role) BatchGet(ctx context.Context, ids []string) ([]*auth.Role, error) {
	var roles []*auth.Role
	return roles, a.DB.WithContext(ctx).Find(&roles, ids).Error
}

func (a *Role) Query(ctx context.Context) ([]*auth.Role, error) {
	var roles []*auth.Role

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

	return roles, tx.Find(&roles).Error
}

func (a *Role) Delete(ctx context.Context, uid string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Where("id = ?", uid).Delete(&auth.Role{})
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Role) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Delete(&auth.Role{}, ids)
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Role) Update(ctx context.Context, role *auth.Role) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Updates(role)
	return executionResult.RowsAffected, executionResult.Error
}
