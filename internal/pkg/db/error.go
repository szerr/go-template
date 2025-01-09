package db

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"go-template/internal/pkg/er"
)

var (
	DuplicateEntry = er.NewErr(204, "DuplicateKey", "duplicate key", er.Info, false) // 重复的 key
)

// ConvertDBError 转换 gorm 和 db 错误 到内部错误
func ConvertDBError(err error) error {
	if err == nil {
		return nil
	}
	var e *mysql.MySQLError
	if errors.As(err, &e) {
		switch e.Number {
		case 1062:
			return DuplicateEntry.WSEF(err).WithMsg(e.Message)
		}
	}
	return er.WSEF(err)
}
