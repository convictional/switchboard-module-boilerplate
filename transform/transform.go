package transform

import (
	"convictional.com/switchboard/env"
	"convictional.com/switchboard/extract"
	"convictional.com/switchboard/load"
	"convictional.com/switchboard/logging"
	"convictional.com/switchboard/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kr/pretty"
	"go.uber.org/zap"
	"strings"
)

// Transform returns processed flagged, the updated model, and error.
// Products that have already been processed will be returned as false.
func Transform(event models.TriggerEvent) error {
	logger := logging.GetLogger()
	if event.Batch {
		logger.Info("Running based on batch event")
		_, _ = ProcessBatchEvent(logger, event)
	}
	logger.Info("Running based on single event")
	var product models.Product
	err := json.Unmarshal(event.Payload, &product)
	if err != nil {
		return err
	}
	updated, productUpdatePayload, err := ProcessSingleProduct(logger, product, event)
	if err != nil {
		return err
	}

	if !updated {
		return nil
	}

	if env.DoLoad() {
		// Marshal into payload
		err = load.LoadSingleProductToConvictionalAPI(product, productUpdatePayload, event)
		if err != nil {
			logger.Error("failed to load product", zap.Error(err))
			return err
		}
	} else {
		logger.Info("Load has not been set")
	}
	return nil
}

func ProcessBatchEvent(logger *zap.Logger, event models.TriggerEvent) (updatedProducts []models.UpdateProduct, err error) {
	products, err := extract.Multiple(event)
	if err != nil {
		logger.Error("failed to get products from extract layer", zap.Error(err))
		return nil, err
	}
	// There are two optional options: transform or load
	for _, product := range products {
		changed, updatedProduct, err := ProcessSingleProduct(logger, product, event)
		if err != nil {
			return nil, err
		}
		if changed {
			updatedProducts = append(updatedProducts, updatedProduct)
		}
	}
	return updatedProducts, nil
}

func ProcessSingleProduct(logger *zap.Logger, product models.Product, event models.TriggerEvent) (changed bool, productUpdate models.UpdateProduct, err error) {
	var updatedProduct = product
	if env.DoTransform() {
		logger.Debug(fmt.Sprintf("Tranforming [%s]", product.ID))
		_, updatedProduct, err = ManipulateProduct(product)
	} else {
		logger.Info("Transform has not been set")
		return false, models.UpdateProduct{}, errors.New("transformer disabled")
	}
	changed, productUpdate = CreateUpdateProductPayload(product, updatedProduct)
	return changed, productUpdate, nil
}

func ManipulateProduct(product models.Product) (processed bool, updated models.Product, err error) {
	// Insert your custom code!
	product.Title = strings.ToUpper(product.Title) // TODO - For demo
	return true, product, nil
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
