package models

import "time"

// https://mholt.github.io/json-to-go/
type BuyerProductResponse struct {
	Data  Product          `json:"data"`
	Error []interface{} `json:"error"`
}

type BuyerProductsResponse struct {
	Data  []Product          `json:"data"`
	HasMore bool `json:"hasMore"`
	Next string `json:"next"`
	Error []interface{} `json:"error"`
}

type SellerProductResponse struct {
	Data  Product          `json:"data"`
	Error []interface{} `json:"error"`
}

type SellerProductsResponse struct {
	Data  []Product          `json:"data"`
	HasMore bool `json:"hasMore"`
	Next string `json:"next"`
	Error []interface{} `json:"error"`
}


type Dimensions struct {
	Length int    `json:"length"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Units  string `json:"units"`
}
type BuyerVariant struct {

}
type Images struct {
	ID         string      `json:"_id"`
	Src        string      `json:"src"`
	Position   int         `json:"position"`
	VariantIds []string `json:"variantIds"`
}

type Variant struct {
	ID                string      `json:"_id"`
	Title             string      `json:"title"`
	RetailPrice       int         `json:"retailPrice"`
	InventoryQuantity int         `json:"inventory_quantity"`
	SkipCount         bool        `json:"skipCount"`
	Weight            int         `json:"weight"`
	WeightUnits       string      `json:"weightUnits"`
	Dimensions        Dimensions  `json:"dimensions"`
	Sku               string      `json:"sku"`
	Barcode           string      `json:"barcode"`
	BarcodeType       string      `json:"barcodeType"`
	Code              string      `json:"code"`
	Attributes        interface{} `json:"attributes"`
	Metafields        interface{} `json:"Metafields"`
	VariantID                int64       `json:"id"`
	Option1           string      `json:"option1"`
	Option2           string      `json:"option2"`
	Option3           string      `json:"option3"`

	InventoryAmount int        `json:"inventoryAmount"`
	RetailCurrency  string     `json:"retailCurrency"`
	BasePrice       int        `json:"basePrice"`
	BaseCurrency    string     `json:"baseCurrency"`
	Options         []Options  `json:"options"`
}
type Options struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	Type     string `json:"type"`
}

type GoogleProductCategory struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}



type Product struct {
	ID                    string                `json:"_id"`
	Code                  string                `json:"code"`
	Active                bool                  `json:"active"`
	BodyHTML              string                `json:"bodyHtml"`
	Images                []Images              `json:"images"`
	Tags                  []string              `json:"tags"`
	Title                 string                `json:"title"`
	Vendor                string                `json:"vendor"`
	Variants              []Variant            `json:"variants"`
	Options               []Options             `json:"options"`
	GoogleProductCategory GoogleProductCategory `json:"googleProductCategory"`
	DelistedUpdated       time.Time             `json:"delistedUpdated"`
	Created               time.Time             `json:"created"`
	Updated               time.Time             `json:"updated"`
	CompanyObjectID       string                `json:"companyObjectId"`
	Attributes            map[string]interface{}           `json:"attributes"`
	CompanyID             string                `json:"companyId"`
	Brand                 string                `json:"brand"`
	Description           string                `json:"description"`
	SellerReference       string                `json:"sellerReference"`
	OptionNames           []string              `json:"optionNames"`
}

