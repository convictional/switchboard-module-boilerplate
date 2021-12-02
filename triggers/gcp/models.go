package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"switchboard-module-boilerplate/models"
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
	var product models.Product

	err := json.Unmarshal(b.Data, &product)
	if err != nil {
		return models.TriggerEvent{}, fmt.Errorf("failed to parse body into product :: %s", err.Error())
	}

	return models.TriggerEvent{
		ID:      b.MessageID,
		Batch:   false,
		Product: &product,
	}, nil
}

func (r *HTTPWebRequest) ConvertHTTPToTriggerEvent() (models.TriggerEvent, error) {
	var body models.ConvictionalWebhookEvent

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()

	err := json.Unmarshal([]byte(s), &body)

	if err != nil {
		return models.TriggerEvent{}, fmt.Errorf("failed to parse body into a webhook event :: %s", err.Error())
	}

	return models.TriggerEvent{
		ID:      body.Id,
		Batch:   false,
		Product: &body.Data,
	}, nil
}
