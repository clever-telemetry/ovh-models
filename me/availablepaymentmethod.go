package me

// GENERATED SDK for me API

// Available payment methods
type AvailablePaymentMethod struct {
	Icon                        Icon            `json:"icon,omitempty"`
	Integration                 IntegrationType `json:"integration,omitempty"`
	Oneshot                     bool            `json:"oneshot,omitempty"`
	PaymentType                 string          `json:"paymentType,omitempty"`
	Registerable                bool            `json:"registerable,omitempty"`
	RegisterableWithTransaction bool            `json:"registerableWithTransaction,omitempty"`
}
