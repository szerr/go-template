package auth

import (
	"go-template/internal/pkg/config"
	"go-template/internal/pkg/er"
	"golang.org/x/crypto/bcrypt"
)

type IAuth interface {
	CompareHashAndPassword(hashedPassword, password string) error
	GenerateFromPassword(pwd string) (string, error)
}

func NewAuth(c *config.Config) IAuth {
	return &Auth{
		cost: &c.App.HashCost,
	}
}

type Auth struct {
	// hash 成本
	cost *int
}

// CompareHashAndPassword 校验 hash 与 密码 是否匹配
func (a *Auth) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateFromPassword 生成 密码 hash
func (a *Auth) GenerateFromPassword(pwd string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(pwd), *a.cost)
	return string(p), er.WSEF(err)
}
