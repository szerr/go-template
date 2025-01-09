package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-template/api/engine/middleware"
	v1 "go-template/api/routers/v1"
	"go-template/internal/global"
	"go-template/internal/pkg/auth"
	"go-template/internal/pkg/config"
	"go.uber.org/zap"
	"net/http"
)

type IRegister interface {
	Register(*gin.RouterGroup)
}

type RoutersV1 struct {
	RegisterLi []IRegister
}

func NewRoutersV1(
	AccessPointRouter *v1.UserRouter,
) *RoutersV1 {
	return &RoutersV1{
		[]IRegister{
			AccessPointRouter,
		},
	}
}

func (a *RoutersV1) Register(gAPI *gin.RouterGroup) {
	v1Group := gAPI.Group("v1")
	for _, r := range a.RegisterLi {
		r.Register(v1Group)
	}
}

func NewEngineBase(register *RoutersV1, c *config.Config, log *zap.Logger, _ *global.Global) EngineBase {
	return EngineBase{
		register: register,
		logLevel: &c.Log.Level,
		httpAddr: c.Http.Addr(),
		log:      log,
	}
}

// EngineBase 运行 gin.Engine
type EngineBase struct {
	register   *RoutersV1
	logLevel   *config.LogLevel
	httpAddr   string
	log        *zap.Logger
	middleware []gin.HandlerFunc
}

func (e *EngineBase) Use(middleware ...gin.HandlerFunc) {
	e.middleware = append(e.middleware, middleware...)
}

// Run 运行服务
func (e *EngineBase) Run(options ...func(*gin.Engine)) error {
	gin.SetMode(gin.DebugMode)
	switch *e.logLevel {
	case config.DebugLevel:
		gin.SetMode(gin.TestMode)
	case config.InfoLevel:
		gin.SetMode(gin.DebugMode)
	case config.WarnLevel, config.ErrLevel:
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(middleware.ErrorHandler(e.log))
	engine.Use(e.middleware...)
	for _, o := range options {
		o(engine)
	}
	gAPI := engine.Group("/api/")
	e.register.Register(gAPI)
	e.log.Info("bind: " + e.httpAddr)
	return engine.Run(e.httpAddr)
}

type EngineDoc struct {
	EngineBase
}

func (e *EngineDoc) Run() error {
	return e.EngineBase.Run(func(e *gin.Engine) {
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		e.GET("/", func(c *gin.Context) { c.Redirect(http.StatusFound, "/swagger/index.html") })
	})
}

type EngineAuth struct {
	EngineBase
	Jwt auth.IJWT
}

func (e *EngineAuth) Run() error {
	e.EngineBase.Use(middleware.JwtAuth(e.Jwt))
	return e.EngineBase.Run()
}

var EngineProviderSet = wire.NewSet(
	NewRoutersV1,
	NewEngineBase,
	wire.Struct(new(EngineDoc), "*"),
	wire.Struct(new(EngineAuth), "*"),
)
