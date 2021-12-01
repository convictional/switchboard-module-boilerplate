package load

import (
	"fmt"
	"go.uber.org/zap"
	"switchboard-module-boilerplate/env"
	"switchboard-module-boilerplate/logging"
	"switchboard-module-boilerplate/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// PublishToSQS :: Sends a message to a SQS.
// https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/sqs/SendMessage/SendMessage.go
func PublishToSQS(productAsBytes []byte, event models.TriggerEvent) error {
	logger := logging.GetLogger()

	// Setup AWS Session
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		logger.Error("failed to setup AWS session", zap.Error(err))
		return err
	}

	//
	svc := sqs.New(awsSession)

	sqsQueue := env.LoadSQS()
	logger.Info(fmt.Sprintf("Found queue :: [%s]", sqsQueue))

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		//MessageAttributes: map[string]*sqs.MessageAttributeValue{},
		MessageBody: aws.String(string(productAsBytes)),
		QueueUrl:   &sqsQueue,
	})

	if err != nil {
		logger.Error("failed to send SQS message", zap.Error(err))
		return err
	}
	return nil
}