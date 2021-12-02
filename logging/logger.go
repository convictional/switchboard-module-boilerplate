package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"convictional.com/switchboard/env"
)

func GetLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	if env.Debug() {
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	} else {
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Failed to setup logger :: %+v", err)
	}
	return logger
}