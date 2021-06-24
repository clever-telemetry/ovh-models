package me

// GENERATED SDK for me API

// Missing description
type BalanceDetails struct {
	Amount         Price              `json:"amount,omitempty"`
	BalanceSubName *string            `json:"balanceSubName,omitempty"`
	Expiring       []ExpiringMovement `json:"expiring,omitempty"`
	ServiceId      *int64             `json:"serviceId,omitempty"`
}
