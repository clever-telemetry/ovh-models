package me

// GENERATED SDK for me API

// Representation of a generic product
type GenericProductDefinition struct {
	PlanCode    string                  `json:"planCode,omitempty"`
	Prices      []GenericProductPricing `json:"prices,omitempty"`
	ProductName string                  `json:"productName,omitempty"`
	ProductType GenericProductTypeEnum  `json:"productType,omitempty"`
}
