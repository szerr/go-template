package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/internal/pkg/er"
	"go.uber.org/zap"
	"net/http"
)

func ErrorHandler(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if v := recover(); v != nil {
				// 输出详细的堆栈信息
				err := er.Internal.WithStackSkip(2).WithZapField(zap.Any("recover", v))
				log.Error(fmt.Sprint(v), zap.String("Stack", err.StackString()))
				c.JSON(http.StatusOK, gin.H{
					"code":    err.Code(),
					"message": err.Msg(),
					"data":    nil,
				})
			}
		}()
		c.Next()
		if len(c.Errors) != 0 {
			for _, e := range c.Errors {
				if e != nil {
					log.Error(e.Error(), zap.Error(e),
						zap.String("url", c.Request.URL.String()),
						zap.String("method", c.Request.Method),
					)
				}
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    er.Unknown.Code(),
				"message": er.Unknown.Msg(),
				"data":    nil,
			})
		}
	}
}
