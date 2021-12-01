package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"switchboard-module-boilerplate/env"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest is the entry point for AWS Lambda
// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
func HandleRequest(ctx context.Context, awsEvent AWSTriggerEvent) {
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

	logger.Debug(fmt.Sprintf("AWS Events :: %+v", awsEvent))

	// Convert event to be platform-agnostic
	event, err := awsEvent.ConvertToTriggerEvent()
	if err != nil {
		logger.Error("Failed to convert trigger event", zap.Error(err))
		return
	}

	service := NewService(logger) // TODO - Move shared drive
	service.Run(event)
}

