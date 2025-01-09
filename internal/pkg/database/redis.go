package database

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-template/internal/pkg/config"
	"go.uber.org/zap"
)

func NewRedis(conf *config.Config, log *zap.Logger) (*redis.Client, func(), error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	return rdb, func() {
		if err := rdb.Close(); err != nil {
			log.Error("d.RedisDB.Close() error", zap.Error(err))
		}
	}, nil
}
