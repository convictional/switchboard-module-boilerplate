package load
import (
	"bytes"
	"convictional.com/switchboard/logging"
	"convictional.com/switchboard/models"
	"convictional.com/switchboard/outbound_http"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"strings"
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
	Headers map[string]string
}

// PublishToAPI will take a config and send the payload to the URL in the config using the send method specified
func PublishToAPI(config APIPublishConfig, event models.TriggerEvent) error {
	logger := logging.GetLogger()
	logger.Debug(fmt.Sprintf("Request body :: %s", string(config.Payload)))
	fmt.Printf("string(config.Payload) :: %+v\n", string(config.Payload))

	rl := rate.NewLimiter(rate.Every(1*time.Second), 5) // 5 request every 1 seconds
	c := outbound_http.NewClient(rl)
	req, _ := http.NewRequest(config.Method, config.URL, bytes.NewReader(config.Payload))
	for headerKey, headerValue := range config.Headers {
		req.Header.Set(headerKey, headerValue)
	}
	resp, err := c.Do(req)
	if err != nil {
		logger.Fatal("Error dispatching HTTP loader request", zap.Error(err))
		return err
	}
	responseText, err := getReader(resp.Body)
	if err != nil {
		logger.Error("failed to read response text", zap.Error(err))
	}
	if resp.StatusCode != http.StatusOK {
		logger.Error("Request delivery failed", zap.Error(err), zap.Int("Status code", resp.StatusCode), zap.String("response", responseText))
		return errors.New("request failed")
	}
	return nil
}

func getReader(r io.Reader) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}