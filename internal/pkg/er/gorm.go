package er

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// 常用异常捕获
func IsGormRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// 常用异常检查，注意底层异常都应该包裹 WithStack，带上异常栈
// CheckRowsAffected 检查 rowsAffected 是否符合数量要求，顺便检查 error
func CheckRowsAffected(db *gorm.DB, rowsAffected int64) IShellError {
	if db.Error != nil {
		return Unknown.WithErr(db.Error).WithStack()
	}
	if db.RowsAffected != rowsAffected {
		return RowsAffectedErr.WithMsgf("should be %d , but the result is %d.", rowsAffected, db.RowsAffected).WithStack()
	}
	return nil
}

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
	return WSEF(err)
}
