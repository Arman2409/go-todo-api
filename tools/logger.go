// logger/logger.go
package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

const (
	LogFilePath = "./logs/logs.txt"
)

func InitLogger() {
	logDir := filepath.Dir(LogFilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   LogFilePath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	zap.ReplaceGlobals(Logger)
}

func LogWithObject(
	message string,
	object interface{},
) {
	objString := fmt.Sprintf("%+v", object)

	Logger.Info(message + ": " + objString)
}
