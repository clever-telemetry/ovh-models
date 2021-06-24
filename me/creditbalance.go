package me

// GENERATED SDK for me API

// Credit balance applied on an Order
type CreditBalance struct {
	Amount      Price  `json:"amount,omitempty"`
	BalanceName string `json:"balanceName,omitempty"`
}
