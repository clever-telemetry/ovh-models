package me

// GENERATED SDK for me API

// All fields needed for a payment mean integration
type AvailablePaymentMeanField struct {
	Key     string                   `json:"key,omitempty"`
	Options *[]string                `json:"options,omitempty"`
	Type    PaymentMeanFieldTypeEnum `json:"type,omitempty"`
	Value   *string                  `json:"value,omitempty"`
}
