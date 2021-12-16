package load

import (
	"convictional.com/switchboard/models"
	"encoding/json"
	"fmt"
)

func LoadSingleProductToConvictionalAPI(product models.Product, update models.UpdateProduct, event models.TriggerEvent) error {
	fmt.Printf("product :: %+v\n", product)
	bytes, err := json.Marshal(update)
	if err != nil {
		return err
	}
	return UpdateProduct(bytes, product.ID, event)
}

func Multiple(products []models.Product, event models.TriggerEvent) error {
	return nil
}
