package logger

import (
	"go-template/internal/pkg/config"
	"go.uber.org/zap"
)

func NewLogger(conf *config.Config) (*zap.Logger, func(), error) {
	cl := conf.Log
	level := zap.DebugLevel
	err := level.UnmarshalText([]byte(cl.Level))
	if err != nil {
		return nil, nil, err
	}
	EncoderConfig := zap.NewDevelopmentEncoderConfig()
	switch conf.Log.EncoderMode {
	case "production":
		EncoderConfig = zap.NewProductionEncoderConfig()
	}
	logger := zap.Must(zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         cl.Encoding,
		OutputPaths:      cl.OutputPaths,
		ErrorOutputPaths: cl.ErrorOutputPaths,
		EncoderConfig:    EncoderConfig,
	}.Build())

	return logger, func() { logger.Sync() }, nil
}

type T struct {
	Level            string   `json:"level"`
	Encoding         string   `json:"encoding"`
	OutputPaths      []string `json:"outputPaths"`
	ErrorOutputPaths []string `json:"errorOutputPaths"`
	InitialFields    struct {
		Foo string `json:"foo"`
	} `json:"initialFields"`
	EncoderConfig struct {
		MessageKey   string `json:"messageKey"`
		LevelKey     string `json:"levelKey"`
		LevelEncoder string `json:"levelEncoder"`
	} `json:"encoderConfig"`
}
