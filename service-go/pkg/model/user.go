package model

import (
	"context"
	"sync"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/ncraft-io/armory/go/pkg/armory/auth"
)

var user *User
var userOnce sync.Once

type User struct {
	DB *db.DB
}

func GetUserModel() *User {
	userOnce.Do(func() {
		user = NewUser()
	})

	return user
}

func NewUser() *User {
	t := &User{DB: GetDB()}
	if !t.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&auth.User{}) {
		if err := d.AutoMigrate(&auth.User{}); err != nil {
			logs.ErrLog("Init User model err: ", err)
			panic(err)
		}
	}
	return t
}

func (a *User) Create(ctx context.Context, users ...*auth.User) (int64, error) {
	usersLen := len(users)
	var executionResult *gorm.DB

	if usersLen == 0 {
		return 0, nil
	} else if usersLen == 1 {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(users[0])
	} else {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(users, len(users))
	}

	return executionResult.RowsAffected, executionResult.Error
}

func (a *User) Get(ctx context.Context, uid string) (*auth.User, error) {
	user := &auth.User{}
	return user, a.DB.WithContext(ctx).First(user, "id = ? or name = ? or email_address = ?", uid, uid, uid).Error
}

func (a *User) BatchGet(ctx context.Context, ids []string) ([]*auth.User, error) {
	var users []*auth.User
	return users, a.DB.WithContext(ctx).Find(&users, ids).Error
}

func (a *User) Query(ctx context.Context) ([]*auth.User, error) {
	var users []*auth.User

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

	return users, tx.Find(&users).Error
}

func (a *User) Delete(ctx context.Context, uid string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Where("id = ?", uid).Delete(&auth.User{})
	return executionResult.RowsAffected, executionResult.Error
}

func (a *User) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Delete(&auth.User{}, ids)
	return executionResult.RowsAffected, executionResult.Error
}

func (a *User) Update(ctx context.Context, user *auth.User) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Updates(user)
	return executionResult.RowsAffected, executionResult.Error
}
