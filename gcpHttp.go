package gcp_trigger

import (
	"convictional.com/switchboard/models"
	"fmt"
	"log"
	"net/http"

	"convictional.com/switchboard/env"
	"convictional.com/switchboard/triggers"
	"convictional.com/switchboard/triggers/gcp"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CVWebhookTrigger is an HTTP Cloud Function with a request parameter.
func CVWebhookTrigger(w http.ResponseWriter, r *http.Request) {

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

	logger.Debug(fmt.Sprintf("GCP Events :: %+v", r))

	// Convert event to be platform-agnostic
	event, err := gcp.ConvertHTTPToTriggerEvent(r)
	if err != nil {
		logger.Error("Failed to convert trigger event", zap.Error(err))
		return
	}

	service := triggers.NewService(logger) // TODO - Move shared drive
	service.Run(event)
}

// GenericHTTPTrigger is an HTTP Cloud Function with a request parameter.
func GenericHTTPTrigger(w http.ResponseWriter, r *http.Request) {

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

	logger.Debug(fmt.Sprintf("GCP Events :: %+v", r))

	service := triggers.NewService(logger) // TODO - Move shared drive
	var bodyBytes []byte
	_, err = r.Body.Read(bodyBytes)
	if err != nil {
		log.Fatalf("Failed to read request body :: %+v", err)
		return
	}
	service.Run(models.TriggerEvent{
		Payload: bodyBytes,
	})
}
