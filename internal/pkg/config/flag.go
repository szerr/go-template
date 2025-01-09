package config

import "flag"

func Flag() *string {
	return flag.String("c", "config.yaml", "path to config file")
}

// FlagParse 获取命令行参数中的配置文件路径
func FlagParse() string {
	confPath := flag.String("c", "config.yaml", "path to config file")
	flag.Parse()
	return *confPath
}

// 初始化配置文件
func InitConfig() (*Config, error) {
	confPath := FlagParse()
	return NewConfig(confPath)
}
