package middleware

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/domain"
	"go-template/internal/global"
	"go-template/internal/pkg/auth"
	"go-template/internal/pkg/er"
	"net/http"
)

func JwtAuth(jwt auth.IJWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-auth-token")
		errResp := domain.Base{
			Code: er.PermissionDenied.Code(),
			Msg:  er.PermissionDenied.Msg(),
		}
		if token == "" {
			c.JSON(http.StatusForbidden, errResp)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusForbidden, errResp)
			c.Abort()
			return
		}

		c.Set(global.ContextClaimsKey, claims)
		c.Next()
	}
}
