package main

import (
	"context"
	"switchboard-module-boilerplate/env"
	"switchboard-module-boilerplate/extract"
	"switchboard-module-boilerplate/load"
	"switchboard-module-boilerplate/models"
	"switchboard-module-boilerplate/transform"
)

// HandleRequest is the entry point for AWS Lambda
// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
func HandleRequest(ctx context.Context, awsEvent AWSTriggerEvent) {
	event, err := awsEvent.ConvertToTriggerEvent()
	if err != nil {
		// TODO - Add logging
		return
	}
	if event.Batch {
		ProcessBatchEvent(event)
	}

	// Load
	if event.Product != nil {
		// TODO - Get single?
	}

	ProcessSingleProduct(*event.Product, event)
}

func ProcessBatchEvent(event models.TriggerEvent) {
	products, err := extract.Multiple(event)
	if err != nil {
		// TODO - Add logging
		return
	}

	// There are two optional options: transform or load
	for _, product := range products {
		ProcessSingleProduct(product, event)
	}
}

func ProcessSingleProduct(product models.Product, event models.TriggerEvent) {
	var err error
	if env.DoTransform() {
		product, err = transform.Transform(product)
		if err != nil {
			// TODO - Add logging
			return
		}
	}

	if env.DoLoad() {
		err = load.Single(product, event)
		if err != nil {
			// TODO - Add logging
		}
	}
}