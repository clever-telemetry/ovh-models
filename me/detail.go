package me

// GENERATED SDK for me API

// Element of consumption for resource
type Detail struct {
	Price     Price   `json:"price,omitempty"`
	Quantity  int64   `json:"quantity,omitempty"`
	Unique_id *string `json:"unique_id,omitempty"`
}
