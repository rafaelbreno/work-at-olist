package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Log    *zap.Logger
	Config zap.Config
}

func init() {
	var l Logger

	l.
		setConfig().
		log()

	defer l.Log.Info("Starting Application")
}

func (l *Logger) setConfig() *Logger {
	l.Config = zap.NewProductionConfig()

	encConfig := zap.NewProductionEncoderConfig()
	encConfig.TimeKey = "timestamp"
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encConfig.StacktraceKey = ""
	l.Config.EncoderConfig = encConfig

	return l
}

func (l *Logger) log() *Logger {

	log, err := l.Config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	l.Log = log

	return l
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
