package main

import (
	"context"
	"fmt"
	"log"
	"switchboard-module-boilerplate/env"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
// triggerPubSub consumes a Pub/Sub message.
func triggerPubSub(ctx context.Context, gcpPubSub GCPRecord) error {
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

	logger.Debug(fmt.Sprintf("AWS Events :: %+v", gcpPubSub))

	// Convert event to be platform-agnostic
	event, err := gcpPubSub.ConvertToTriggerEvent()
	if err != nil {
		logger.Error("Failed to convert trigger event", zap.Error(err))
		return err
	}

	service := NewService(logger) // TODO - Move shared drive
	service.Run(event)

	return err
}
