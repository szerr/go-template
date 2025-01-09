package main

import (
	"go-template/internal/pkg/config"
	"os"
	"path"
)

// 生成配置文件
func main() {
	confPath := config.FlagParse()
	err := os.MkdirAll(path.Dir(confPath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = config.GenConfig(confPath)
	if err != nil {
		panic(err)
	}
}
