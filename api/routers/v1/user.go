package v1

import (
	"github.com/gin-gonic/gin"
	"go-template/api/handle"
)

func NewUserRouter(routerGroup *gin.RouterGroup) *UserRouter {
	h := new(UserRouter)
	h.Register(routerGroup)
	return h
}

type UserRouter struct {
	UserHandle *handle.UserHandle
}

func (a *UserRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/user")
	g.POST("create", a.UserHandle.Create)
	g.POST("update", a.UserHandle.Update)
	g.GET("delete/:id", a.UserHandle.Delete)
	g.POST("list", a.UserHandle.List)
	g.GET("retrieve/:ap_code", a.UserHandle.Retrieve)
}
