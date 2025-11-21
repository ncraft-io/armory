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

var policy *Policy
var policyOnce sync.Once

type Policy struct {
	DB *db.DB
}

func GetPolicyModel() *Policy {
	policyOnce.Do(func() {
		policy = NewPolicy()
	})

	return policy
}

func NewPolicy() *Policy {
	t := &Policy{DB: GetDB()}
	if !t.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&auth.Policy{}) {
		if err := d.AutoMigrate(&auth.Policy{}); err != nil {
			logs.ErrLog("Init Policy model err: ", err)
			panic(err)
		}
	}
	return t
}

func (a *Policy) Create(ctx context.Context, policys ...*auth.Policy) (int64, error) {
	policysLen := len(policys)
	var executionResult *gorm.DB

	if policysLen == 0 {
		return 0, nil
	} else if policysLen == 1 {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(policys[0])
	} else {
		executionResult = a.DB.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(policys, len(policys))
	}

	return executionResult.RowsAffected, executionResult.Error
}

func (a *Policy) Get(ctx context.Context, uid string) (*auth.Policy, error) {
	policy := &auth.Policy{}
	return policy, a.DB.WithContext(ctx).First(policy, "id = ?", uid).Error
}

func (a *Policy) BatchGet(ctx context.Context, ids []string) ([]*auth.Policy, error) {
	var policys []*auth.Policy
	return policys, a.DB.WithContext(ctx).Find(&policys, ids).Error
}

func (a *Policy) Query(ctx context.Context) ([]*auth.Policy, error) {
	var policys []*auth.Policy

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

	return policys, tx.Find(&policys).Error
}

func (a *Policy) Delete(ctx context.Context, uid string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Where("id = ?", uid).Delete(&auth.Policy{})
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Policy) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Delete(&auth.Policy{}, ids)
	return executionResult.RowsAffected, executionResult.Error
}

func (a *Policy) Update(ctx context.Context, policy *auth.Policy) (int64, error) {
	executionResult := a.DB.WithContext(ctx).Updates(policy)
	return executionResult.RowsAffected, executionResult.Error
}
