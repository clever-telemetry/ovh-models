package me

// GENERATED SDK for me API

// A validation required to add a payment mean credit card
type CreditCardValidation struct {
	Id             int64              `json:"id,omitempty"`
	SubmitUrl      *string            `json:"submitUrl,omitempty"`
	Url            string             `json:"url,omitempty"`
	ValidationType ValidationTypeEnum `json:"validationType,omitempty"`
}
