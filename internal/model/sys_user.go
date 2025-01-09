package model

import (
	"go-template/internal/global"
	"go-template/internal/pkg/er"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	ID        uint64  `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`
	NickName  string  `gorm:"column:nick_name;type:varchar(64);not null;comment:昵称" json:"nick_name"`                                          // 暱称
	Email     *string `gorm:"column:email;type:varchar(320);uniqueIndex:sys_user_unique_1,priority:1;comment:邮箱" json:"email"`                 // 邮箱
	UserName  string  `gorm:"column:user_name;type:varchar(100);not null;uniqueIndex:sys_user_unique,priority:1;comment:用户名" json:"user_name"` // 用户名
	Password  string  `gorm:"column:password;type:varchar(100);not null;comment:密码" json:"password"`                                           // 密码
	GroupName *string `gorm:"column:group_name;type:varchar(100);comment:角色名" json:"group_name"`                                               // 角色名
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}

// SetPwd 设置密码
func (m *SysUser) SetPwd(pwd string) error {
	p, err := global.Auth.GenerateFromPassword(pwd)
	if err != nil {
		return er.WSEF(err)
	}
	m.Password = p
	return nil
}

// VerifyPwd 校验密码
func (m *SysUser) VerifyPwd(pwd string) error {
	err := global.Auth.CompareHashAndPassword(m.Password, pwd)
	return er.WSEF(err)
}
