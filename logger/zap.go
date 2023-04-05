package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.Logger
}

func New() *zapLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return &zapLogger{logger: logger}
}

func (z *zapLogger) Print(message string, properties map[string]any) {
	z.logger.Info(message, zap.Any("properties", properties))
}

func (z *zapLogger) Error(err error, properties map[string]any) {
	z.logger.Error(err.Error(), zap.Any("properties", properties))
}

func (z *zapLogger) Fatal(err error, properties map[string]any) {
	z.logger.Fatal(err.Error(), zap.Any("properties", properties))
}
