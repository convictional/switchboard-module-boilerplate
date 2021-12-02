package transform

import (
	"strings"
	"convictional.com/switchboard/models"
)

func Transform(product models.Product) (models.Product, error) {
	// Insert your custom code!
	product.Title = strings.ToUpper(product.Title) // TODO - For demo
	return product, nil
}
