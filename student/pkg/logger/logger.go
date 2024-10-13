package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"student/pkg/setting"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggingSetting) *LoggerZap {
	var level zapcore.Level
	switch config.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		MaxBackups: config.MaxBackups,
	}
	encoder := getEncoderLog()
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout)),
		level,
	)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DebugLevel))}
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.TimeKey = "time"
	return zapcore.NewJSONEncoder(encoderConfig)
}
