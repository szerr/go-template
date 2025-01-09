package biz

import (
	"context"
	"go-template/internal/domain"
	"go-template/internal/model"
	"go-template/internal/pkg/auth"
	"go-template/internal/pkg/er"
	"go.uber.org/zap"
)

type IUserRepo interface {
	Create(ctx context.Context, cp *model.SysUser) error
	Update(ctx context.Context, cp *model.SysUser) error
	Delete(ctx context.Context, id uint64) error
	PageInfo(ctx context.Context, req *domain.UserListRequest) ([]*model.SysUser, int64, error)
	Retrieve(ctx context.Context, id uint64) (*model.SysUser, error)
	GetByUserName(ctx context.Context, userName string) (*model.SysUser, error)
}

type UserBiz struct {
	Log      *zap.Logger
	UserRepo IUserRepo
	Jwt      auth.IJWT
}

// SigIn 校验密码并生成 Jwt
func (s *UserBiz) SigIn(ctx context.Context, userName, password string) (string, error) {
	user, err := s.UserRepo.GetByUserName(ctx, userName)
	if err != nil {
		return "", err
	}
	err = user.VerifyPwd(password)
	if err != nil {
		return "", er.WrongUserNameOrPassword
	}
	return s.Jwt.GenToken(user.ID)
}

func (s *UserBiz) SigOut(ctx context.Context, username, password string) error {
	return nil
}

func (s *UserBiz) Create(ctx context.Context, cp *model.SysUser, pwd string) error {
	err := cp.SetPwd(pwd)
	if err != nil {
		return err
	}
	return s.UserRepo.Create(ctx, cp)
}

func (s *UserBiz) Update(ctx context.Context, cp *model.SysUser) error {
	return s.UserRepo.Update(ctx, cp)
}

func (s *UserBiz) Delete(ctx context.Context, id uint64) error {
	return s.UserRepo.Delete(ctx, id)
}

func (s *UserBiz) List(ctx context.Context, req *domain.UserListRequest) (*domain.PageData, error) {
	li, count, err := s.UserRepo.PageInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return domain.FillInPageResponseData(req.PageSearch, count, li), nil
}

func (s *UserBiz) Retrieve(ctx context.Context, id uint64) (*model.SysUser, error) {
	return s.UserRepo.Retrieve(ctx, id)
}
