package models

type UpdateProduct struct {
	Active *bool `json:"active,omitempty"`
	BodyHTML              *string                `json:"bodyHtml,omitempty"`
	Images *[]Images `json:"images,omitempty"`
	Tags                  *[]string              `json:"tags,omitempty"`
	Title                 *string                `json:"title,omitempty"`
	Vendor                *string                `json:"vendor,omitempty"`
	Variants              *[]Variant            `json:"variants,omitempty"`
	Options               *[]Options             `json:"options,omitempty"`
	GoogleProductCategory *GoogleProductCategory `json:"googleProductCategory,omitempty"`
	Attributes            *map[string]interface{}           `json:"attributes,omitempty"`
}
