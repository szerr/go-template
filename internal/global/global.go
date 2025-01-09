package global

import "go-template/internal/pkg/auth"

type Global struct{}

var Auth auth.IAuth

// InitGlobal 初始化全局变量
func InitGlobal(auth auth.IAuth) *Global {
	Auth = auth
	return &Global{}
}
