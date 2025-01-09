package data

import (
	"context"
	"go-template/internal/domain"
	"go-template/internal/model"
)

func NewUserRepo(base *BaseRepo) *UserRepo {
	return &UserRepo{
		BaseRepo: base,
	}
}

type UserRepo struct {
	*BaseRepo
}

func (r *UserRepo) Create(ctx context.Context, cp *model.SysUser) error {
	return r.db.Create(cp).Error()
}

func (r *UserRepo) Update(ctx context.Context, cp *model.SysUser) error {
	return r.db.Omit("created_at").Save(cp).Error()
}

func (r *UserRepo) Delete(ctx context.Context, id uint64) error {
	return r.db.Model(new(model.SysUser)).Where("id=?", id).Error()
}

func (r *UserRepo) PageInfo(ctx context.Context, req *domain.UserListRequest) ([]*model.SysUser, int64, error) {
	li := []*model.SysUser{}
	count, err := r.BaseRepo.PageInfo(r.db, req, &li)
	return li, count, err
}

func (r *UserRepo) List(ctx context.Context, req *domain.UserListRequest) ([]*model.SysUser, error) {
	li := []*model.SysUser{}
	return li, r.IntoPages(req).Find(&li).Error()
}

func (r *UserRepo) Count(ctx context.Context, req *domain.UserListRequest) (count int64, err error) {
	return count, r.db.Count(&count).Error()
}

func (r *UserRepo) Retrieve(ctx context.Context, id uint64) (*model.SysUser, error) {
	m := new(model.SysUser)
	return m, r.db.Where("id=?", id).First(m).Error()
}

func (r *UserRepo) GetByUserName(ctx context.Context, userName string) (*model.SysUser, error) {
	m := new(model.SysUser)
	return m, r.db.Where("user_name =?", userName).First(m).Error()
}
