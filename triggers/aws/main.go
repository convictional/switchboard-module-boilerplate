package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
	"switchboard-module-boilerplate/logging"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest is the entry point for AWS Lambda
// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
func HandleRequest(ctx context.Context, awsEvent AWSTriggerEvent) {
	logger := logging.GetLogger()

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

