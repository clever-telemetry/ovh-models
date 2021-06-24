package me

// GENERATED SDK for me API

// Payment methods
type PaymentMethods struct {
	Available  []AvailablePaymentMethod `json:"available,omitempty"`
	Registered []int64                  `json:"registered,omitempty"`
}
