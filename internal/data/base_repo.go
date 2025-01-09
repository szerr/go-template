package data

import (
	"github.com/redis/go-redis/v9"
	"go-template/internal/pkg/db"
)

func NewBaseRepo(db db.IDB, rdb *redis.Client) *BaseRepo {
	return &BaseRepo{
		db:  db,
		rdb: rdb,
	}
}

type BaseRepo struct {
	db  db.IDB
	rdb *redis.Client
}

type PageReq interface {
	GetPageNum() int
	GetPageSize() int
}

// IntoPages 实现分页逻辑
func (r *BaseRepo) IntoPages(req PageReq) db.IDB {
	pageNum := req.GetPageNum()
	pageSize := req.GetPageSize()
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 || pageSize > 1000 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	return r.db.Offset(offset).Limit(pageNum)
}

// PageInfo 返回 list 和 count，需要传入一个经过 业务 条件筛选的 db
func (r *BaseRepo) PageInfo(db db.IDB, req PageReq, v any) (count int64, err error) {
	err = db.Model(v).Count(&count).Error()
	if err != nil {
		return
	}
	b := &BaseRepo{db: db}
	return count, b.IntoPages(req).Find(v).Error()
}

type TimeReq interface {
	GetDateType() string
	GetStartTime() int64
	GetEndTime() int64
}

// TimeFilter 时间筛选
func (r *BaseRepo) TimeFilter(req TimeReq) db.IDB {
	field := "created_at"
	switch req.GetDateType() {
	case "create_time":
		field = "created_at"
	case "update_time":
		field = "updated_at"
	}
	db := r.db
	if req.GetStartTime() > 0 {
		db = r.db.Where("? >= ?", field, req.GetStartTime())
	}
	if req.GetEndTime() > 0 {
		db = db.Where("? <= ?", field, req.GetEndTime())
	}
	return db
}

// PageTimeFilter 处理分页和时间筛选
func (r *BaseRepo) PageTimeFilter(db db.IDB, req any, v any) (count int64, err error) {
	d := db
	if e, ok := req.(TimeReq); ok {
		b := &BaseRepo{db: db}
		d = b.TimeFilter(e)
	}
	if e, ok := req.(PageReq); ok {
		return r.PageInfo(d, e, v)
	}
	return count, d.Count(&count).Error()
}
