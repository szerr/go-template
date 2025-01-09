package pkg

import (
	"github.com/google/wire"
	"go-template/internal/global"
	"go-template/internal/pkg/auth"
	"go-template/internal/pkg/config"
	"go-template/internal/pkg/database"
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/logger"
	"go-template/internal/pkg/snowflake"
)

var PkgProviderSet = wire.NewSet(
	config.InitConfig,
	logger.NewLogger,
	snowflake.Init,
	database.NewDB,
	database.NewRedis,
	auth.NewJWT,
	auth.NewAuth,
	db.NewDB,
	global.InitGlobal,
)
