package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap() (logger *zap.Logger, err error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err = zapConfig.Build(zap.AddCallerSkip(1))
	return
}
