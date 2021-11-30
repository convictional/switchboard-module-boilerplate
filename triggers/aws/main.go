package main

import (
	"context"
	"go.uber.org/zap"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest is the entry point for AWS Lambda
// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
func HandleRequest(ctx context.Context, awsEvent AWSTriggerEvent) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Convert event to be platform-agnostic
	event, err := awsEvent.ConvertToTriggerEvent()
	if err != nil {
		logger.Error("Failed to convert trigger event", zap.Error(err))
		return
	}

	service := NewService(logger) // TODO - Move shared drive
	service.Run(event)
}

