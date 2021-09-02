package global

import (
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func initLogger() {
	// set log rotation
	fileWriter, err := rotateLogs.New(
		path.Join("logs", "%Y-%m-%d.log"),
		rotateLogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(fmt.Errorf("failed to set log rotation: %s \n", err))
	}
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	// set level
	var level zapcore.Level
	switch Config.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	default:
		level = zapcore.InfoLevel
	}
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		writeSyncer,
		level,
	)
	// add stack trace in debug
	if level == zapcore.DebugLevel {
		Logger = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		Logger = zap.New(core)
	}
}
