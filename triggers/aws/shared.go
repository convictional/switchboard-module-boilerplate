package main

import (
	"go.uber.org/zap"
	"switchboard-module-boilerplate/env"
	"switchboard-module-boilerplate/extract"
	"switchboard-module-boilerplate/load"
	"switchboard-module-boilerplate/models"
	"switchboard-module-boilerplate/transform"
)

type Service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) Service {
	return Service{
		logger: logger,
	}
}

func (s *Service) Run(event models.TriggerEvent) {
	if event.Batch {
		s.ProcessBatchEvent(event)
	}

	// Load
	if event.Product != nil {
		// TODO - Get single?
	}

	s.ProcessSingleProduct(*event.Product, event)
}

func (s *Service) ProcessBatchEvent(event models.TriggerEvent) {
	products, err := extract.Multiple(event)
	if err != nil {
		// TODO - Add logging
		return
	}

	// There are two optional options: transform or load
	for _, product := range products {
		s.ProcessSingleProduct(product, event)
	}
}

func (s *Service) ProcessSingleProduct(product models.Product, event models.TriggerEvent) {
	var err error
	if env.DoTransform() {
		product, err = transform.Transform(product)
		if err != nil {
			s.logger.Error("failed to transform product", zap.Error(err))
			return
		}
	}

	if env.DoLoad() {
		err = load.Single(product, event)
		if err != nil {
			s.logger.Error("failed to load product", zap.Error(err))
		}
	}
}