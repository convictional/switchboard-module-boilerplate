package extract

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"io/ioutil"
	"net/http"
	"switchboard-module-boilerplate/env"
	"switchboard-module-boilerplate/logging"
	"switchboard-module-boilerplate/models"
	"switchboard-module-boilerplate/outbound_http"
	"time"
)

func GetProductsFromAPI(event models.TriggerEvent) ([]models.Product, error) {
	logger := logging.GetLogger()

	rl := rate.NewLimiter(rate.Every(1*time.Second), 5) // 5 request every 1 seconds
	c := outbound_http.NewClient(rl)
	hasMore := true
	page := 1
	results := []models.Product{}
	for hasMore {
		reqURL := fmt.Sprintf("https://api.convictional.com/products?page=%d&limit=25", page)
		if env.IsBuyer() {
			reqURL = fmt.Sprintf("https://api.convictional.com/buyer/products?page=%d&limit=25", page)
		}
		req, _ := http.NewRequest("GET", reqURL, nil)
		resp, err := c.Do(req)
		if err != nil {
			logger.Error("failed to make an HTTP call", zap.Error(err))
			return []models.Product{}, err
		}
		if resp.StatusCode == 429 {
			// This should not be hittable
			logger.Error("calls are being rate limited", zap.Error(err))
			return []models.Product{}, errors.New("rate limit hit")
		}
		if resp.StatusCode != 200 {
			logger.Error("received non 200 response from get products", zap.Int("statusCode", resp.StatusCode))
			return []models.Product{}, errors.New("get products received non 200 response")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Error("failed to parse the response body", zap.Int("statusCode", resp.StatusCode))
			return []models.Product{}, err
		}

		if env.IsBuyer() {
			var productsResponse models.BuyerProductsResponse
			err = json.Unmarshal(body, &productsResponse)
			if err != nil {
				logger.Error("failed to unmarshal product response", zap.Int("statusCode", resp.StatusCode))
				return []models.Product{}, err
			}

			if productsResponse.HasMore {
				hasMore = false
			}
			results = append(results, productsResponse.Data...)
		} else {
			var productsResponse models.SellerProductsResponse
			err = json.Unmarshal(body, &productsResponse)
			if err != nil {
				logger.Error("failed to unmarshal product response", zap.Int("statusCode", resp.StatusCode))
				return []models.Product{}, err
			}

			if productsResponse.HasMore {
				hasMore = false
			}
			results = append(results, productsResponse.Data...)
		}
	}

	return results, nil
}


func GetProductFromAPI(productID string) (models.Product, error) {
	logger := logging.GetLogger()

	rl := rate.NewLimiter(rate.Every(1*time.Second), 5) // 5 request every 1 seconds
	c := outbound_http.NewClient(rl)

	reqURL := fmt.Sprintf("https://api.convictional.com/products/%s", productID)
	if env.IsBuyer() {
		reqURL = fmt.Sprintf("https://api.convictional.com/buyer/products/%s", productID)
	}
	req, _ := http.NewRequest("GET", reqURL, nil)
	resp, err := c.Do(req)
	if err != nil {
		logger.Error("failed to make an HTTP call", zap.Error(err))
		return models.Product{}, err
	}
	if resp.StatusCode == 429 {
		// This should not be hittable
		logger.Error("calls are being rate limited", zap.Error(err))
		return models.Product{}, errors.New("rate limit hit")
	}
	if resp.StatusCode != 200 {
		logger.Error("received non 200 response from get products", zap.Int("statusCode", resp.StatusCode))
		return models.Product{}, errors.New("get products received non 200 response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to parse the response body", zap.Int("statusCode", resp.StatusCode))
		return models.Product{}, err
	}

	if env.IsBuyer() {
		var productsResponse models.BuyerProductResponse
		err = json.Unmarshal(body, &productsResponse)
		if err != nil {
			logger.Error("failed to unmarshal product response", zap.Int("statusCode", resp.StatusCode))
			return models.Product{}, err
		}

		return productsResponse.Data, nil
	}

	var productsResponse models.SellerProductResponse
	err = json.Unmarshal(body, &productsResponse)
	if err != nil {
		logger.Error("failed to unmarshal product response", zap.Int("statusCode", resp.StatusCode))
		return models.Product{}, err
	}

	return productsResponse.Data, nil
}