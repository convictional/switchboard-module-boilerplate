package main

import (
	"encoding/json"
	"fmt"
	"switchboard-module-boilerplate/models"
)

type GCPRecord struct {
	MessageID   string `json:"messageId"`
	Data        []byte `json:"data"`
	Attributes  string `json:"attributes"`
	OrderingKey string `json:"orderingKey"`
}

func (b *GCPRecord) ConvertToTriggerEvent() (models.TriggerEvent, error) {

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
