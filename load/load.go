package load

import (
	"convictional.com/switchboard/env"
	"convictional.com/switchboard/models"
	"errors"
	"fmt"
)

const (
	LoadMethodConvictionalAPI = "convictional_api"
)

func Single(product models.Product, updatedProduct models.Product, event models.TriggerEvent) error {
	fmt.Printf("product :: %+v\n", product)
	switch env.LoadMethod() {
	case LoadMethodConvictionalAPI:
		return UpdateProduct(product, updatedProduct, event)
	default:
		return errors.New("invalid load method")
	}
	return nil
}
func Multiple( products []models.Product, event models.TriggerEvent) error {
	return nil
}
