package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"switchboard-module-boilerplate/models"
)

const (
	SourceSQS = "aws:sqs"
)

type AWSTriggerEvent struct {
	Records []AWSRecord `json:"records"`
}

func (b *AWSTriggerEvent) ConvertToTriggerEvent() (models.TriggerEvent, error) {
	if len(b.Records) != 1 {
		return models.TriggerEvent{}, errors.New("no records found")
	}

	var product models.Product
	err := json.Unmarshal([]byte(b.Records[0].Body), &product)
	if err != nil {
		return models.TriggerEvent{}, fmt.Errorf("failed to parse body into product :: %s", err.Error())
	}

	return models.TriggerEvent{
		ID: b.Records[0].MessageID,
		Batch: b.Records[0].Batch(),
		Product: &product,
	}, nil
}

type AWSRecord struct {
	MessageID string `json:"messageId"`
	Body string `json:"body"`
	EventSource string `json:"eventSource"`
	EventSourceARN string `json:"eventSourceARN"`
}

func (r *AWSRecord) Batch() bool {
	return r.EventSource == SourceSQS
}