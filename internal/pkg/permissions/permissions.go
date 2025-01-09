package permissions

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func NewCasbinEnforcer(db *gorm.DB) (*casbin.Enforcer, error) {
	//  设置 casbin 规则
	m, err := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && r.act == p.act
	`)
	if err != nil {
		return nil, err
	}
	gormadapter.TurnOffAutoMigrate(db) // 关闭表自动迁移
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	// Load the policy from DB.
	// e.LoadPolicy()

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	//e.SavePolicy()
	return e, err
}
