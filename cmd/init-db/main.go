package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

// 初始化数据
func main() {
	casbinEnforcer, cleanup, err := wireApp()
	if err != nil {
		log.Panic(err)
	}
	defer cleanup()
	// 添加管理员角色
	_, err = casbinEnforcer.AddPolicy("admin", "/api/*", "GET")
	if err != nil {
		log.Panic(err)
	}
	// 添加管理员用户
	_, err = casbinEnforcer.AddGroupingPolicy("admin", "admin")
	if err != nil {
		log.Panic(err)
	}
	// 验证权限
	//ok, err := casbinEnforcer.Enforce("admin", "/api/a", "GET")
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println(ok, 1111111111)
}

func newApp(casbinEnforcer *casbin.Enforcer) *casbin.Enforcer {
	return casbinEnforcer
}
