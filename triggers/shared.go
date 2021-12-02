package triggers

import (
	"convictional.com/switchboard/env"
	"convictional.com/switchboard/extract"
	"convictional.com/switchboard/load"
	"convictional.com/switchboard/models"
	"convictional.com/switchboard/transform"
	"fmt"
	"go.uber.org/zap"
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
	s.logger.Info("Test Logger")
	if event.Batch {
		s.logger.Info("Running based on batch event")
		s.ProcessBatchEvent(event)
		return
	}
	s.logger.Info("Running based on single event")

	// Load
	if event.Product != nil {
		// TODO - Get single?
	}


	s.ProcessSingleProduct(*event.Product, event)
}

func (s *Service) ProcessBatchEvent(event models.TriggerEvent) {
	products, err := extract.Multiple(event)
	if err != nil {
		s.logger.Error("failed to get products from extract layer", zap.Error(err))
		return
	}

	// There are two optional options: transform or load
	for _, product := range products {
		s.ProcessSingleProduct(product, event)
	}
}

func (s *Service) ProcessSingleProduct(product models.Product, event models.TriggerEvent) {
	var err error
	var processed bool
	var updatedProduct models.Product
	if env.DoTransform() {
		s.logger.Debug(fmt.Sprintf("Tranforming [%s]", product.ID))
		processed, updatedProduct, err = transform.Transform(product)
		if !processed {
			s.logger.Info(fmt.Sprintf("product [%s] has already been processed", product.ID))
			return
		}
	} else {
		s.logger.Info("Transform has not been set")
	}

	if env.DoLoad() {
		s.logger.Debug(fmt.Sprintf("Loading [%s]", product.ID))
		err = load.Single(product, updatedProduct, event)
		if err != nil {
			s.logger.Error("failed to load product", zap.Error(err))
		}
	} else {
		s.logger.Info("Load has not been set")
	}
}