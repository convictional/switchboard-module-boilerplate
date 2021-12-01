package load

import (
	"bytes"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"net/http"
	"switchboard-module-boilerplate/logging"
	"switchboard-module-boilerplate/models"
	"switchboard-module-boilerplate/outbound_http"
	"time"
)

const (
	POST = "POST"
	PATCH = "PATCH"
)

type APIPublishConfig struct {
	Payload []byte
	Method string
	URL string
}

// PublishToAPI will take a config and send the payload to the URL in the config using the send method specified
func PublishToAPI(config APIPublishConfig, event models.TriggerEvent) error {
	logger := logging.GetLogger()
	rl := rate.NewLimiter(rate.Every(1*time.Second), 5) // 5 request every 1 seconds
	c := outbound_http.NewClient(rl)
	req, _ := http.NewRequest(config.Method, config.URL, bytes.NewReader(config.Payload))
	resp, err := c.Do(req)
	if err != nil {
		logger.Fatal("Error dispatching HTTP loader request", zap.Error(err))
		return err
	}
	if resp.StatusCode != http.StatusOK {
		logger.Error("Request delivery failed", zap.Error(err))
		return errors.New("request failed")
	}
	return nil
}
