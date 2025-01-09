package database

import (
	"fmt"
	"go-template/internal/pkg/config"
	"go-template/internal/pkg/er"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(conf *config.Config, log *zap.Logger) (*gorm.DB, func(), error) {
	var logModel = logger.Info
	switch conf.Log.Level {
	case config.DebugLevel, config.InfoLevel:
		logModel = logger.Info
	case config.WarnLevel:
		logModel = logger.Warn
	case config.ErrLevel:
		logModel = logger.Error
	default:
		return nil, nil, er.ConfigError.WSF(zap.String("level", string(conf.Log.Level)))
	}

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB.Username, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		//SkipDefaultTransaction: true, //禁用事务
		Logger: logger.Default.LogMode(logModel),
	})
	return db, func() {
		db, err := db.DB()
		if err != nil {
			log.Error("d.DB.DB() error", zap.Error(err))
		}
		if err := db.Close(); err != nil {
			log.Error("d.DB.Close() error", zap.Error(err))
		}
	}, err
}
