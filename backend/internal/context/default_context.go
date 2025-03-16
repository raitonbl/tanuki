package context

import (
	"github.com/raitonbl/tanuki/internal/config"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type DefaultContext struct {
	logger        *zap.Logger
	configuration config.Config
}

func NewContext(cfg config.Config) Context {
	return &DefaultContext{
		configuration: cfg,
		logger:        createZapLogger(cfg),
	}
}

func createZapLogger(cfg config.Config) *zap.Logger {
	zapLogLevel := zapcore.InfoLevel
	if cfg.LogLevel == config.DebugLogLevel {
		zapLogLevel = zapcore.DebugLevel
	}
	fields := []zap.Field{
		zap.String("service", cfg.Service),
		zap.String("environment", cfg.Environment),
	}
	if cfg.Solution != "" {
		fields = append(fields, zap.String("solution", cfg.Solution))
	}
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zapLogLevel)
	logger := zap.New(core, zap.AddCaller())
	if len(fields) > 0 {
		logger = logger.With(fields...)
	}
	return logger
}

func (d DefaultContext) Logger() *zap.Logger {
	return d.logger
}

func (d DefaultContext) Configuration() config.Config {
	return d.configuration
}
