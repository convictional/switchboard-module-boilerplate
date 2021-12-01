package extract

import "switchboard-module-boilerplate/models"

func Single(event models.TriggerEvent) (models.Product, error) {
	return models.Product{}, nil
}

func Multiple(event models.TriggerEvent) ([]models.Product, error) {
	return []models.Product{}, nil
}
