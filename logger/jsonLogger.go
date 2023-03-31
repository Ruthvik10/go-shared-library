package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}
func Print(message string, properties map[string]any) {
	log.Info(message, zap.Any("properties", properties))
}

func Error(err error, properties map[string]any) {
	log.Error(err.Error(), zap.Any("properties", properties))
}

func Fatal(err error, properties map[string]any) {
	log.Fatal(err.Error(), zap.Any("properties", properties))
}
