package transform

import (
	"strings"
	"switchboard-module-boilerplate/models"
)

func Transform(product models.Product) (models.Product, error) {
	// Insert your custom code!
	product.Title = strings.ToUpper(product.Title) // TODO - For demo
	return product, nil
}
