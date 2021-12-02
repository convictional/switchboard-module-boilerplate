package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"switchboard-module-boilerplate/models"
)

type GCPPubSubRecord struct {
	MessageID   string `json:"messageId"`
	Data        []byte `json:"data"`
	Attributes  string `json:"attributes"`
	OrderingKey string `json:"orderingKey"`
}

type GCPWebhookEvent struct {
	Id              string         `json:"_id"`
	Type            string         `json:"type"`
	Created         string         `json:"created"`
	CompanyObjectId string         `json:"companyObjectId`
	Data            models.Product `json:"data"`
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

func ConvertHTTPToTriggerEvent(r *http.Request) (models.TriggerEvent, error) {
	var body GCPWebhookEvent

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
