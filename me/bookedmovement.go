package me

// GENERATED SDK for me API

// Missing description
type BookedMovement struct {
	Amount         Price   `json:"amount,omitempty"`
	BalanceSubName *string `json:"balanceSubName,omitempty"`
	OrderId        int64   `json:"orderId,omitempty"`
}
