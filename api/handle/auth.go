package handle

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go-template/internal/biz"
	"go-template/internal/domain"
)

type AuthHandle struct {
	*BaseHandle
	UserBiz *biz.UserBiz

	CasbinService *casbin.Enforcer
}

// @Tags Auth
// @Summary SigIn
// @Param body domain.SigIn true "Request body"
// @Success 200 {string} string
// @Router /auth/sig_in [post]
func (h *AuthHandle) SigIn(c *gin.Context) {
	req := new(domain.SigIn)
	if h.BindJSON(c, req) {
		return
	}
	data, err := h.UserBiz.SigIn(c.Request.Context(), req.User, req.Password)
	h.FeedBack(c, data, err)
}

// @Tags Auth
// @Security SigOut
// @Param body any true "Request body"
// @Success 200 {object} model.SysUser
// @Router /auth/sig_out [post]
func (h *AuthHandle) SigOut(c *gin.Context) {
	h.Success(c, nil)
}
