package main

import (
	"fmt"
	"log"
	"net/http"
	"switchboard-module-boilerplate/env"
	"switchboard-module-boilerplate/triggers/shared"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// httpTriggerEvent is an HTTP Cloud Function with a request parameter.
func httpTriggerEvent(w http.ResponseWriter, r HTTPWebRequest) {

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
		return
	}

	logger.Debug(fmt.Sprintf("AWS Events :: %+v", r))

	// Convert event to be platform-agnostic
	event, err := r.ConvertHTTPToTriggerEvent(r)
	if err != nil {
		logger.Error("Failed to convert trigger event", zap.Error(err))
		return
	}

	service := shared.NewService(logger) // TODO - Move shared drive
	service.Run(event)

	return
}
