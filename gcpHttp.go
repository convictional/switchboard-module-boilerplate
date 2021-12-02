package gcp_trigger

import (
	"fmt"
	"log"
	"net/http"

	"convictional.com/switchboard/env"
	"convictional.com/switchboard/triggers/gcp"
	"convictional.com/switchboard/triggers/shared"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// httpTriggerEvent is an HTTP Cloud Function with a request parameter.
func HttpTriggerEvent(w http.ResponseWriter, r *http.Request) error {

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
		return err
	}

	logger.Debug(fmt.Sprintf("GCP Events :: %+v", r))

	// Convert event to be platform-agnostic
	event, err := gcp.ConvertHTTPToTriggerEvent(r)
	if err != nil {
		logger.Error("Failed to convert trigger event", zap.Error(err))
		return err
	}

	service := shared.NewService(logger) // TODO - Move shared drive
	service.Run(event)
	return nil
}
