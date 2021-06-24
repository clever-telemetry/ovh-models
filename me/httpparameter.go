package me

// GENERATED SDK for me API

// Parameter to give to a payment page
type HttpParameter struct {
	Choice *[]HttpParameterChoice `json:"choice,omitempty"`
	Name   string                 `json:"name,omitempty"`
	Value  *string                `json:"value,omitempty"`
}
