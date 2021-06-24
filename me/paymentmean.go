package me

// GENERATED SDK for me API

// All data needed to use a payment mean
type PaymentMean struct {
	Fee        float64         `json:"fee,omitempty"`
	HtmlForm   *string         `json:"htmlForm,omitempty"`
	HttpMethod string          `json:"httpMethod,omitempty"`
	Logo       *string         `json:"logo,omitempty"`
	Parameters []HttpParameter `json:"parameters,omitempty"`
	SubType    *string         `json:"subType,omitempty"`
	Url        string          `json:"url,omitempty"`
}
