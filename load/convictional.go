package load

import (
	"convictional.com/switchboard/env"
	"convictional.com/switchboard/logging"
	"convictional.com/switchboard/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kr/pretty"
)

func UpdateProduct(product models.Product, updatedProduct models.Product, event models.TriggerEvent) error {
	logger := logging.GetLogger()

	diff, updateProductPayload := CreateUpdateProductPayload(product, updatedProduct)
	if !diff {
		// Two models are the same, no need to update
		return nil
	}

	// Marshal into payload
	productPayloadAsBytes, err := json.Marshal(updateProductPayload)
	if err != nil {
		return nil
	}

	url := fmt.Sprintf("%s/products/%s", env.ConvictionalAPIURL(), product.ID)
	if env.IsBuyer() {
		return errors.New("no endpoint exists yet")
	}

	logger.Debug(fmt.Sprintf("Calling with payload :: %+v", updateProductPayload))

	return PublishToAPI(APIPublishConfig{
		Payload: productPayloadAsBytes,
		Method: "PATCH",
		URL: url,
		Headers: map[string]string{
			"Authorization": env.ConvictionalAPIKeyForLoad(),
			"Content-Type": "application/json",
		},
	}, event)
}

// CreateUpdateProductPayload compares the two objects, and creates an update. This being done in
func CreateUpdateProductPayload(product models.Product, updatedProduct models.Product) (bool, models.UpdateProduct) {
	updateProductPayload := models.UpdateProduct{}
	if product.Title != updatedProduct.Title {
		updateProductPayload.Title = &updatedProduct.Title
	}
	if product.Description != updatedProduct.Description {
		updateProductPayload.BodyHTML = &updatedProduct.Description
	}
	if product.Active != updatedProduct.Active {
		updateProductPayload.Active = &updatedProduct.Active
	}
	if !match(product.Tags, updatedProduct.Tags) {
		updateProductPayload.Tags = &updatedProduct.Tags
	}
	if !match(product.Options, updatedProduct.Options) {
		updateProductPayload.Options = &updatedProduct.Options
	}
	if product.GoogleProductCategory.Name != updatedProduct.GoogleProductCategory.Name {
		updateProductPayload.GoogleProductCategory.Name = updatedProduct.GoogleProductCategory.Name
	}
	if product.GoogleProductCategory.Code != updatedProduct.GoogleProductCategory.Code {
		updateProductPayload.GoogleProductCategory.Code = updatedProduct.GoogleProductCategory.Code
	}
	if !match(product.Images, updatedProduct.Images) {
		updateProductPayload.Images = &updatedProduct.Images
	}
	if !match(product.Variants, updatedProduct.Variants) {
		updateProductPayload.Variants = &updatedProduct.Variants
	}
	if !match(product.Attributes, updatedProduct.Attributes) {
		updateProductPayload.Attributes = &updatedProduct.Attributes
	}

	return true, updateProductPayload
}

func match(obj1 interface{}, obj2 interface{}) bool {
	fieldsThatDiffer := pretty.Diff(obj1, obj2)
	return 0 == len(fieldsThatDiffer)
}