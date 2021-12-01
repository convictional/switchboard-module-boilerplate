package load

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"go.uber.org/zap"
	"switchboard-module-boilerplate/env"
	"switchboard-module-boilerplate/logging"
	"switchboard-module-boilerplate/models"
)

// PublishToPubSub :: Sends a message to a GCP Pub/Sub.
// https://github.com/GoogleCloudPlatform/golang-samples/blob/HEAD/appengine_flexible/pubsub/pubsub.go
func PublishToPubSub(bytePayload []byte, _ models.TriggerEvent) error {
	logger := logging.GetLogger()
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, env.GCP_PROJECT())
	if err != nil {
		logger.Fatal("Failed creating new pub/sub client", zap.Error(err))
		return err
	}
	defer client.Close()

	topicName := env.GCP_LOAD_TOPIC_NAME()
	topic := client.Topic(topicName)

	// Create the topic if it doesn't exist.
	exists, err := topic.Exists(ctx)
	if err != nil {
		logger.Fatal("Failed checking if pub/sub topic exists", zap.Error(err))
		return err
	}
	if !exists {
		logger.Info(fmt.Sprintf("Topic %v doesn't exist - creating it", topicName))
		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			logger.Fatal("Failed creating new pub/sub topic", zap.Error(err))
			return err
		}
	}

	msg := &pubsub.Message{
		Data: bytePayload,
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		logger.Error("Could not publish message", zap.Error(err))
		return err
	}
	logger.Info("Message published")
	return nil
}