package models

type TriggerEvent struct {
	ID      string
	Batch   bool
	Product *Product
}

type ConvictionalWebhookEvent struct {
	Id              string  `json:"_id"`
	Type            string  `json:"type"`
	Created         string  `json:"created"`
	CompanyObjectId string  `json:"companyObjectId`
	Data            Product `json:"data"`
}
