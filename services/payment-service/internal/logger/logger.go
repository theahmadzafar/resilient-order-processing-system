package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel string
}

type Logger struct {
	Zap      *zap.Logger
	LogLevel zapcore.Level
}

func NewLogger(cfg *Config) (*Logger, error) {
	zapCfg := zap.NewDevelopmentConfig()

	level, err := zap.ParseAtomicLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	zapCfg.Level = level

	zapLogger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(zapLogger)

	return &Logger{Zap: zapLogger, LogLevel: level.Level()}, nil
}
