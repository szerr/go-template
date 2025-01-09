package v1

import (
	"github.com/gin-gonic/gin"
	"go-template/api/handle"
)

func NewAuthRouter(routerGroup *gin.RouterGroup) *UserRouter {
	h := new(UserRouter)
	h.Register(routerGroup)
	return h
}

type AuthRouter struct {
	AuthHandle *handle.AuthHandle
}

func (a *AuthRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/auth")
	g.POST("sig_in", a.AuthHandle.SigIn)
	g.POST("sig_out", a.AuthHandle.SigOut)
}
