package extract

import "switchboard-module-boilerplate/models"

func Single(event models.TriggerEvent) models.Product {
	return models.Product{}
}

func Multiple(event models.TriggerEvent) []models.Product {
	return []models.Product{}
}
