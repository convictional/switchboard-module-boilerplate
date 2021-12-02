package transform

import (
	"convictional.com/switchboard/models"
	"strings"
)

// Transform returns processed flagged, the updated model, and error.
// Products that have already been processed will be returned as false.
func Transform(product models.Product) (bool, models.Product, error) {
	// Insert your custom code!
	product.Title = strings.ToUpper(product.Title) // TODO - For demo
	return true, product, nil
}
