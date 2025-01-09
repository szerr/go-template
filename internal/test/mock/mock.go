package mock

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go-template/internal/biz"
	"go-template/internal/test/mock/mock-biz"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MockProviderSet is mock providers.
var MockProviderSet = wire.NewSet(
	NewSqlMock,
	NewGormMock,
	NewRedisMock,
	wire.Bind(new(biz.IUserRepo), new(*mock_biz.MockIUserRepo)),
)

type SqlMock struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

func NewSqlMock() (*SqlMock, func(), error) {
	db, mock, err := sqlmock.New()
	return &SqlMock{
		DB:   db,
		Mock: mock,
	}, func() { db.Close() }, err
}

func NewGormMock(mock *SqlMock) (*gorm.DB, error) {
	// 处理 gorm 的初始化步骤，返回 mysql 版本
	mock.Mock.ExpectQuery("SELECT VERSION()").WillReturnRows(
		sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.4.0"),
	)
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mock.DB,
	}), &gorm.Config{})
	return gdb, err
}

type RedisMock struct {
	RDB  *redis.Client
	Mock redismock.ClientMock
}

func NewRedisMock() *RedisMock {
	db, mock := redismock.NewClientMock()
	return &RedisMock{
		RDB:  db,
		Mock: mock,
	}
}

type DataMock struct {
	SqlDB     *sql.DB
	SqlMock   sqlmock.Sqlmock
	DB        *gorm.DB
	RedisDB   *redis.Client
	RedisMock redismock.ClientMock
}

func NewDataMock() (*DataMock, func(), error) {
	sm, sqlMockClose, err := NewSqlMock()
	if err != nil {
		return nil, nil, err
	}
	gormMock, err := NewGormMock(sm)
	if err != nil {
		return nil, nil, err
	}
	rm := NewRedisMock()
	return &DataMock{
		SqlDB:     sm.DB,
		SqlMock:   sm.Mock,
		DB:        gormMock,
		RedisDB:   rm.RDB,
		RedisMock: rm.Mock,
	}, func() { sqlMockClose() }, nil
}
