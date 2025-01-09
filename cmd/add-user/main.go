package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-template/internal/model"
	"log"
)

// 启动服务
func main() {
	userName := flag.String("u", "admin", "user name")
	groupName := flag.String("g", "", "user group")
	pwd := flag.String("p", "", "password")
	userBiz, cleanup, err := wireApp()
	if err != nil {
		log.Panic(err)
	}
	defer cleanup()
	user := new(model.SysUser)
	user.UserName = *userName
	user.GroupName = groupName
	err = userBiz.Create(new(gin.Context), user, *pwd)
	if err != nil {
		log.Panic(err)
	}
}
