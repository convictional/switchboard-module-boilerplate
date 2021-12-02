package extract

import (
	"convictional.com/switchboard/env"
	"convictional.com/switchboard/logging"
	"convictional.com/switchboard/models"
	"errors"
	"fmt"
)

const (
	ExtractMethodConvictionalAPI = "convictional_api"
)

func Multiple(event models.TriggerEvent) ([]models.Product, error) {
	logger := logging.GetLogger()
	switch env.ExtractMethod() {
	case ExtractMethodConvictionalAPI:
		return GetProductsFromAPI(event)
	default:
		logger.Info(fmt.Sprintf("unsupported extract method :: [%s]", env.ExtractMethod()))
		return []models.Product{}, errors.New("invalid extract method")
	}
	return []models.Product{}, nil
}