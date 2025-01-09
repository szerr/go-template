package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"go-template/internal/domain"
	"go-template/internal/pkg/er"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
	"strings"
)

type IBaseHandleErr interface {
	Name() string
	Level() zapcore.Level
	ToMap() map[string]interface{}
	Code() uint32
	Msg() string
	SendToClient() bool // 是否发送给调用者
	Fields() []zapcore.Field
}

type BaseHandle struct {
	Log *zap.Logger
}

func (h BaseHandle) ParamInt64(c *gin.Context, key string, id *int64) bool {
	var err error
	*id, err = strconv.ParseInt(c.Param(key), 10, 64)
	if err != nil {
		h.SendErr(c, er.InvalidArgument.WithErr(err))
		return true
	}
	return false
}

func (h BaseHandle) ParamUint64(c *gin.Context, key string, id *uint64) bool {
	var err error
	*id, err = strconv.ParseUint(c.Param(key), 10, 64)
	if err != nil {
		h.SendErr(c, er.InvalidArgument.WithErr(err))
		return true
	}
	return false
}

// ParamId 获取 Param 中的 id，在 <= 0 时抛出错误
func (h BaseHandle) ParamId(c *gin.Context, key string, id *uint64) bool {
	if h.ParamUint64(c, key, id) {
		return true
	}
	if *id <= 0 {
		h.SendErr(c, er.InvalidArgument.WithMsg("id cannot be 0"))
		return true
	}
	return false
}

// ParamStrId 获取 Param 中的 string id，在 空值 时抛出错误
func (h BaseHandle) ParamStrId(c *gin.Context, key string, p *string) bool {
	*p = c.Param(key)
	if *p == "" {
		h.SendErr(c, er.InvalidArgument)
		return true
	}
	return false
}

// Copy fromValue to toValue by copier
func (h BaseHandle) Copy(c *gin.Context, toValue interface{}, fromValue interface{}) bool {
	return h.HeadErr(c, copier.Copy(toValue, fromValue))
}

func (h BaseHandle) SendErr(c *gin.Context, err error) {
	var ok bool
	var errInf IBaseHandleErr
	if errInf, ok = err.(IBaseHandleErr); !ok {
		errInf = er.Unknown.WithErr(err)
	}
	h.Log.Log(errInf.Level(), errInf.Msg(), errInf.Fields()...)
	c.JSON(http.StatusBadRequest, domain.Base{
		Code: errInf.Code(),
		Msg:  errInf.Msg(),
	})
}

// HeadErr 处理错误并记录log，如果返回 true，代表发生了错误并被处理，Handle层应该 return，不该继续执行。如果返回 false，说明没有错误，继续执行。
func (h BaseHandle) HeadErr(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	h.SendErr(c, err)
	return true
}

// BindJSON 解析 json，如果返回 true，代表发生了错误并被处理，Handle层应该 return，不该继续执行。如果返回 false，代表没有错误，继续执行。
func (h BaseHandle) BindJSON(c *gin.Context, ptr interface{}) bool {
	if err := c.ShouldBindJSON(&ptr); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			var errList []string
			for _, e := range errs {
				errList = append(errList, e.Error())
			}
			return h.HeadErr(c, er.InvalidArgument.WSEF(err).WithMsg(strings.Join(errList, ",")))
		} else {
			return h.HeadErr(c, err)
		}
	}
	return false
}

// Success 返回成功数据
func (h BaseHandle) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, domain.Base{
		Code: er.Ok.Code(),
		Msg:  er.Ok.Msg(),
		Data: data,
	})
}

// FeedBack 同时处理数据和错误，是 HeadErr 与 Success 的胶水方法，存在错误返回 true，成功返回 false
func (h BaseHandle) FeedBack(c *gin.Context, data interface{}, err error) bool {
	if h.HeadErr(c, err) {
		return true
	}
	h.Success(c, data)
	return false
}
