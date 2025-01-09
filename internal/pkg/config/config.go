package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	yamlEncoder "github.com/zwgblue/yaml-encoder"
	"go-template/internal/pkg/er"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"reflect"
)

type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel           = "info"
	WarnLevel           = "warn"
	ErrLevel            = "error"
)

type Http struct {
	Host string `yaml:"host" comment:"" validate:"required"`
	Port int    `yaml:"port" comment:"" validate:"required"`
	Mode string `yaml:"mode" comment:"" validate:"required"`
}

func (c *Http) Init() {
	c.Host = "0.0.0.0"
	c.Port = 8080
	c.Mode = "debug"
}

func (h *Http) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type DB struct {
	Username string `yaml:"user_name" comment:"" validate:"required"`
	Password string `yaml:"password" comment:"" validate:"required"`
	Host     string `yaml:"host" comment:"" validate:"required"`
	Port     int    `yaml:"port" comment:"" validate:"required"`
	DBName   string `yaml:"db_name" comment:"" validate:"required"`
}

func (c *DB) Init() {
	c.Host = "mysql"
	c.Port = 3306
	c.Username = "go-template"
	c.Password = "go-template"
	c.DBName = "go-template"
}

type Log struct {
	Level            LogLevel `json:"level" yaml:"level" comment:"log level, one of: debug, info, error." validate:"oneof=debug info error"`
	Encoding         string   `json:"encoding" yaml:"encoding" comment:"encoding sets the logger's encoding. Valid values are json and console" validate:"oneof=json console"`
	OutputPaths      []string `json:"output_paths" yaml:"output_paths" comment:"" validate:"required"`
	ErrorOutputPaths []string `json:"error_output_paths" yaml:"error_output_paths" comment:"" validate:"required"`
	EncoderMode      string   `json:"encoder_mode" yaml:"encoder_mode" comment:"encoder_mode is set zap EncoderConfig. one of: develop, production" validate:"oneof=develop production"` // "develop" or "production"
}

func (c *Log) Init() {
	c.Level = "debug"
	c.Encoding = "console"
	c.OutputPaths = []string{"stdout"}
	c.ErrorOutputPaths = []string{"stdout"}
	c.EncoderMode = "develop"
}

// jwt配置
type Jwt struct {
	AccessExpire int64  `yaml:"access_expire" comment:"" validate:"required"` // 访问过期时间
	Secret       string `yaml:"secret" comment:"" validate:"required"`        // 签名
	Issuer       string `yaml:"issuer" comment:"" validate:"required"`        // 发行人
}

func (c *Jwt) Init() {
	c.AccessExpire = 86400
	c.Secret = uuid.NewString()
	c.Issuer = "go-template"
}

// redis 配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`

	DB int `yaml:"db"`
}

func (c *Redis) Init() {
	c.Host = "redis"
	c.Port = 6379
	c.Password = ""
	c.DB = 0
}

type App struct {
	NodeId   int64 `yaml:"node_id" comment:"" validate:"required"`
	HashCost int   `yaml:"hash_cost" comment:"hash cost, range 4~31" validate:"required"`
}

func (c *App) Init() {
	c.NodeId = 1
	c.HashCost = 4
}

type Config struct {
	App   App   `comment:"app is configuration node related information" validate:"required"`
	Http  Http  `comment:"" validate:"required"`
	DB    DB    `comment:"" validate:"required"`
	Redis Redis `comment:"" validate:""`
	Log   Log   `comment:"log is configure zap's output format" validate:"required"`
	Jwt   Jwt   `comment:"" validate:"required"`
}

type configItem interface {
	Init()
}

// AutoCompletion 为配置生成默认值
func (c *Config) AutoCompletion() {
	validate := validator.New(validator.WithRequiredStructEnabled())
	for _, i := range []configItem{&c.App, &c.Http, &c.DB, &c.Redis, &c.Log, &c.Jwt} {
		if i == nil {
			v := reflect.ValueOf(i)
			elem := v.Elem()
			newElem := reflect.New(elem.Type()).Elem()
			elem.Set(newElem)
		}
		if validate.Struct(i) != nil {
			i.Init()
		}
	}
}

// readConfig 读取配置文件
func readConfig(configPath string) (*Config, error) {
	conf := new(Config)
	fd, err := os.Open(configPath)
	if err != nil {
		return conf, err
	}
	confData, err := io.ReadAll(fd)
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(confData, conf)
	if err != nil {
		return nil, err
	}
	return conf, err
}

// GenConfig 生成配置文件，如果配置文件已存在，则读取配置，补全配置，删除不支持的配置。
func GenConfig(configPath string) error {
	var conf *Config
	var err error

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		conf = new(Config)
	} else {
		conf, err = readConfig(configPath)
		if err != nil {
			return err
		}
		// 备份配置文件
		srcFile, err := os.Open(configPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(configPath + ".bak")
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}
	// 配置自动补全
	conf.AutoCompletion()

	confData, err := yamlEncoder.NewEncoder(conf, yamlEncoder.WithComments(yamlEncoder.CommentsOnHead)).Encode()
	if err != nil {
		return er.WSEF(err)
	}
	fd, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	n, err := fd.Write(confData)
	if err != nil {
		return err
	}
	if n != len(confData) {
		return io.ErrShortWrite
	}
	return nil
}

// NewConfig 创建 config 对象
func NewConfig(configPath string) (*Config, error) {
	conf, err := readConfig(configPath)
	if err != nil {
		return nil, err
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(conf)
	return conf, er.ConfigError.WSEF(err)
}
