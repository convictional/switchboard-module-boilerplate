package gcp_trigger

import (
	"context"
	"fmt"
	"log"

	"convictional.com/switchboard/env"
	"convictional.com/switchboard/triggers/gcp"
	"convictional.com/switchboard/triggers/shared"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TriggerPubSub consumes a Pub/Sub message.
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
func TriggerPubSub(ctx context.Context, gcpPubSub gcp.GCPPubSubRecord) error {
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

	logger.Debug(fmt.Sprintf("GCP Events :: %+v", gcpPubSub))

	// Convert event to be platform-agnostic
	event, err := gcpPubSub.ConvertPSToTriggerEvent()
	if err != nil {
		logger.Error("Failed to convert trigger event", zap.Error(err))
		return err
	}

	service := shared.NewService(logger)
	service.Run(event)

	return err
}
