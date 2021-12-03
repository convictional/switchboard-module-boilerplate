package gcp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"convictional.com/switchboard/models"
)

type GCPPubSubRecord struct {
	MessageID   string `json:"messageId"`
	Data        []byte `json:"data"`
	Attributes  string `json:"attributes"`
	OrderingKey string `json:"orderingKey"`
}

type HTTPWebRequest struct {
	Body io.ReadCloser
}

func (b *GCPPubSubRecord) ConvertPSToTriggerEvent() (models.TriggerEvent, error) {
	return models.TriggerEvent{
		ID:      b.MessageID,
		Batch:   false,
		Payload: b.Data,
	}, nil
}

func ConvertHTTPToTriggerEvent(r *http.Request) (models.TriggerEvent, error) {
	var body models.ConvictionalWebhookEvent

	defer r.Body.Close()
	s, err := io.ReadAll(r.Body)
	if err != nil {
		return models.TriggerEvent{}, fmt.Errorf("failed to parse body into a byte array :: %s", err.Error())
	}

	errJson := json.Unmarshal(s, &body)

	if errJson != nil {
		return models.TriggerEvent{}, fmt.Errorf("failed to parse body into a webhook event :: %s", err.Error())
	}

	return models.TriggerEvent{
		ID:      body.Id,
		Batch:   false,
		Payload: body.Data,
	}, nil
}
