package me

// GENERATED SDK for me API

// All data needed to use a payment mean
type AvailablePaymentMean struct {
	Fields      *[]AvailablePaymentMeanField `json:"fields,omitempty"`
	Integration PaymentMeanIntegrationEnum   `json:"integration,omitempty"`
	Name        string                       `json:"name,omitempty"`
	Url         *string                      `json:"url,omitempty"`
}
