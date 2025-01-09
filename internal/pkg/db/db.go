package db

import (
	"database/sql"
	"gorm.io/gorm"
)

type IDB interface {
	RowsAffected() int64
	Error() error
	Omit(columns ...string) *DB
	Select(query interface{}, args ...interface{}) *DB
	Model(value any) *DB
	Count(count *int64) *DB
	Where(query interface{}, args ...interface{}) *DB
	Create(value any) *DB
	Delete(value any, conds ...any) *DB
	Save(value any) *DB
	Update(column string, value interface{}) *DB
	Updates(value interface{}) *DB
	Find(dest interface{}, conds ...interface{}) *DB
	First(dest interface{}, conds ...interface{}) *DB
	Take(dest interface{}, conds ...interface{}) *DB
	Last(dest interface{}, conds ...interface{}) *DB
	Not(query interface{}, args ...interface{}) *DB
	Or(query interface{}, args ...interface{}) *DB
	Order(value any) *DB
	Joins(query string, args ...interface{}) *DB
	Scan(dest any) *DB
	Limit(limit int) *DB
	Offset(offset int) *DB
	Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) error
	Begin(opts ...*sql.TxOptions) *DB
	Rollback() *DB
	Commit() *DB
	SavePoint(name string) *DB
	RollbackTo(name string) *DB
}

type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) IDB {
	return &DB{db}
}

// 返回影响的行
func (d *DB) RowsAffected() int64 {
	return d.db.RowsAffected
}

// Error 返回转换后的错误
func (d *DB) Error() error {
	return ConvertDBError(d.db.Error)
}

// Omit 排除
func (d *DB) Omit(columns ...string) *DB {
	return &DB{db: d.db.Omit(columns...)}
}

// Select 选择
func (d *DB) Select(query interface{}, args ...interface{}) *DB {
	return &DB{db: d.db.Select(query, args...)}
}

func (d *DB) Model(value any) *DB {
	return &DB{db: d.db.Model(value)}
}

func (d *DB) Count(count *int64) *DB {
	return &DB{db: d.db.Count(count)}
}

func (d *DB) Where(query interface{}, args ...interface{}) *DB {
	return &DB{db: d.db.Where(query, args)}
}

func (d *DB) Create(value any) *DB {
	return &DB{db: d.db.Create(value)}
}

func (d *DB) Delete(value any, conds ...any) *DB {
	return &DB{db: d.db.Delete(value, conds...)}
}

func (d *DB) Save(value any) *DB {
	return &DB{db: d.db.Save(value)}
}

func (d *DB) Update(column string, value interface{}) *DB {
	return &DB{db: d.db.Update(column, value)}
}

func (d *DB) Updates(value interface{}) *DB {
	return &DB{db: d.db.Updates(value)}
}

func (d *DB) Find(dest interface{}, conds ...interface{}) *DB {
	return &DB{db: d.db.Find(dest, conds...)}
}

// First 按照主键排序的第一条
func (d *DB) First(dest interface{}, conds ...interface{}) *DB {
	return &DB{db: d.db.First(dest, conds...)}
}

// Take 第一条记录
func (d *DB) Take(dest interface{}, conds ...interface{}) *DB {
	return &DB{db: d.db.Take(dest, conds...)}
}

// Last 最后一条记录
func (d *DB) Last(dest interface{}, conds ...interface{}) *DB {
	return &DB{db: d.db.Last(dest, conds...)}
}

func (d *DB) Not(query interface{}, args ...interface{}) *DB {
	return &DB{db: d.db.Not(query, args...)}
}

func (d *DB) Or(query interface{}, args ...interface{}) *DB {
	return &DB{db: d.db.Or(query, args...)}
}

func (d *DB) Order(value any) *DB {
	return &DB{db: d.db.Order(value)}
}

func (d *DB) Joins(query string, args ...interface{}) *DB {
	return &DB{db: d.db.Joins(query, args...)}
}

func (d *DB) Scan(dest any) *DB {
	return &DB{db: d.db.Scan(dest)}
}

func (d *DB) Limit(limit int) *DB {
	return &DB{db: d.db.Limit(limit)}
}

func (d *DB) Offset(offset int) *DB {
	return &DB{db: d.db.Offset(offset)}
}

// Transaction 执行事务
func (d *DB) Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) error {
	return ConvertDBError(
		d.db.Transaction(
			func(tx *gorm.DB) error {
				return fc(&DB{db: tx})
			},
			opts...))
}

// Begin 开始手动事务
func (d *DB) Begin(opts ...*sql.TxOptions) *DB {
	return &DB{db: d.db.Begin(opts...)}
}

// Rollback 回滚事务
func (d *DB) Rollback() *DB {
	return &DB{db: d.db.Rollback()}
}

// Commit 提交事务
func (d *DB) Commit() *DB {
	return &DB{db: d.db.Commit()}
}

// SavePoint 保存事务
func (d *DB) SavePoint(name string) *DB {
	return &DB{db: d.db.SavePoint(name)}
}

// RollbackTo 回滚事务到保存的事务
func (d *DB) RollbackTo(name string) *DB {
	return &DB{db: d.db.RollbackTo(name)}
}
