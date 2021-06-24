package me

// GENERATED SDK for me API

// Information about a Deposit entry
type DepositDetail struct {
	DepositDetailId string `json:"depositDetailId,omitempty"`
	Description     string `json:"description,omitempty"`
	Domain          string `json:"domain,omitempty"`
	Quantity        string `json:"quantity,omitempty"`
	TotalPrice      Price  `json:"totalPrice,omitempty"`
	UnitPrice       Price  `json:"unitPrice,omitempty"`
}
