package me

// GENERATED SDK for me API

// Details about a payment
type Payment struct {
	PaymentDate       string          `json:"paymentDate,omitempty"`
	PaymentIdentifier *string         `json:"paymentIdentifier,omitempty"`
	PaymentType       PaymentMeanEnum `json:"paymentType,omitempty"`
}
