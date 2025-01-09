package v1

import (
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

func (a *TestRouter) Test(c *gin.Context) {
	println(c.Request.URL)
}

func (a *TestRouter) Register(group *gin.RouterGroup) {
	g := group.Group("/test")
	g.GET("", a.Test)
}
