package logger

import (
	"os"

	"github.com/auth_service/pkg/settings"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	//Timestring to ISO8601 format. E.g 1770732229 ->  Tuesday, February 10, 2026 2:03:49 PM (GMT)
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts -> Time
	encodeConfig.TimeKey = "time"

	//from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

// func getWriterSync() zapcore.WriteSyncer {
// 	file, _ := os.OpenFile
// }

func NewLogger(config settings.LogSetting) *LoggerZap {
	logLevel := config.Log_level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename: config.File_log_name,
		MaxSize:  config.Max_size,
		MaxAge:   config.Max_age,
		Compress: config.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}
