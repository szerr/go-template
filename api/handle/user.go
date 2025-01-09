package handle

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go-template/internal/biz"
	"go-template/internal/domain"
	"go-template/internal/model"
)

type UserHandle struct {
	*BaseHandle
	UserBiz       *biz.UserBiz
	CasbinService *casbin.Enforcer
}

// @Tags User
// @Security ApiKeyAuth
// @Summary Create User
// @Param body model.SysUser true "Request body"
// @Success 200 {object} model.SysUser
// @Router /user/create [post]
func (h *UserHandle) Create(c *gin.Context) {
	req := new(model.SysUser)
	if h.BindJSON(c, req) {
		return
	}
	err := h.UserBiz.Create(c.Request.Context(), req, req.Password)
	if h.HeadErr(c, err) {
		return
	}
	h.Success(c, req)
}

// @Tags User
// @Security ApiKeyAuth
// @Summary Update User
// @Param body model.SysUser true "Request body"
// @Success 200 {object} model.SysUser
// @Router /user/update [post]
func (h *UserHandle) Update(c *gin.Context) {
	req := new(model.SysUser)
	if h.BindJSON(c, req) {
		return
	}
	err := h.UserBiz.Update(c.Request.Context(), req)
	h.FeedBack(c, req, err)
}

// @Tags User
// @Security ApiKeyAuth
// @Summary Delete role record by ID
// @Param id path string true "unique id"
// @Success 200 {object} model.SysUser
// @Router /user/delete/{id} [get]
func (h *UserHandle) Delete(c *gin.Context) {
	var id uint64
	if h.ParamId(c, "id", &id) {
		return
	}
	err := h.UserBiz.Delete(c.Request.Context(), id)
	h.FeedBack(c, nil, err)
}

// @Tags User
// @Security ApiKeyAuth
// @Summary List User
// @Param body domain.UserListRequest true "Request body"
// @Success 200 {object} model.SysUser
// @Router /user/list [post]
func (h *UserHandle) List(c *gin.Context) {
	req := new(domain.UserListRequest)
	if h.BindJSON(c, req) {
		return
	}
	data, err := h.UserBiz.List(c.Request.Context(), req)
	h.FeedBack(c, data, err)
}

// @Tags User
// @Security ApiKeyAuth
// @Summary Retrieve User
// @Success 200 {object} model.SysUser
// @Router /user/retrieve/{id} [get]
func (h *UserHandle) Retrieve(c *gin.Context) {
	var id uint64
	if h.ParamId(c, "id", &id) {
		return
	}
	data, err := h.UserBiz.Retrieve(c.Request.Context(), id)
	h.FeedBack(c, data, err)
}
