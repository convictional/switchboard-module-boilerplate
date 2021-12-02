package load

import (
	"fmt"
	"convictional.com/switchboard/models"
)

func Single(product models.Product, event models.TriggerEvent) error {
	fmt.Printf("product :: %+v\n", product)
	return nil
}

func Multiple( products []models.Product, event models.TriggerEvent) error {
	return nil
}
