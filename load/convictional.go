package load

import (
	"convictional.com/switchboard/env"
	"convictional.com/switchboard/logging"
	"convictional.com/switchboard/models"
	"errors"
	"fmt"
)

func UpdateProduct(payload []byte, productID string, event models.TriggerEvent) error {
	logger := logging.GetLogger()
	url := fmt.Sprintf("%s/products/%s", env.ConvictionalAPIURL(), productID)
	if env.IsBuyer() {
		return errors.New("no endpoint exists yet")
	}

	logger.Debug(fmt.Sprintf("Calling with payload :: %+v", payload))

	return PublishToAPI(APIPublishConfig{
		Payload: payload,
		Method:  "PATCH",
		URL:     url,
		Headers: map[string]string{
			"Authorization": env.ConvictionalAPIKeyForLoad(),
			"Content-Type":  "application/json",
		},
	}, event)
}
