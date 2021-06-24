package me

// GENERATED SDK for me API

// Information about a Withdrawal entry
type WithdrawalDetail struct {
	Description        string `json:"description,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Quantity           string `json:"quantity,omitempty"`
	TotalPrice         Price  `json:"totalPrice,omitempty"`
	UnitPrice          Price  `json:"unitPrice,omitempty"`
	WithdrawalDetailId string `json:"withdrawalDetailId,omitempty"`
}
