package me

// GENERATED SDK for me API

// A validation required to add a payment mean
type PaymentMeanValidation struct {
	Id             int64                     `json:"id,omitempty"`
	SubmitUrl      *string                   `json:"submitUrl,omitempty"`
	Url            string                    `json:"url,omitempty"`
	ValidationType PaymentMeanValidationType `json:"validationType,omitempty"`
}
