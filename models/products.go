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


type Options struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Dimensions struct {
	Length int    `json:"length"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Units  string `json:"units"`
}
type Attributes struct {
}
type Variants struct {
	ID              string     `json:"id"`
	Sku             string     `json:"sku"`
	Title           string     `json:"title"`
	InventoryAmount int        `json:"inventoryAmount"`
	RetailPrice     int        `json:"retailPrice"`
	RetailCurrency  string     `json:"retailCurrency"`
	BasePrice       int        `json:"basePrice"`
	BaseCurrency    string     `json:"baseCurrency"`
	Barcode         string     `json:"barcode"`
	BarcodeType     string     `json:"barcodeType"`
	Options         []Options  `json:"options"`
	SkipCount       bool       `json:"skipCount"`
	Weight          float64    `json:"weight"`
	WeightUnits     string     `json:"weightUnits"`
	Dimensions      Dimensions `json:"dimensions"`
	Attributes      Attributes `json:"attributes"`
}
type Images struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}
type GoogleProductCategory struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

type Product struct {
	ID                    string                `json:"id"`
	CompanyID             string                `json:"companyId"`
	Title                 string                `json:"title"`
	Brand                 string                `json:"brand"`
	Description           string                `json:"description"`
	SellerReference       string                `json:"sellerReference"`
	Variants              []Variants            `json:"variants"`
	Tags                  []string              `json:"tags"`
	Images                []Images              `json:"images"`
	OptionNames           []string              `json:"optionNames"`
	GoogleProductCategory GoogleProductCategory `json:"googleProductCategory"`
	Attributes            Attributes            `json:"attributes"`
	Created               time.Time             `json:"created"`
	Updated               time.Time             `json:"updated"`
}