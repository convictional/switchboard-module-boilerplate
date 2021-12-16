package triggers

import (
	"convictional.com/switchboard/models"
	"convictional.com/switchboard/transform"
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
	transform.Transform(event)
}
