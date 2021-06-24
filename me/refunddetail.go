package me

// GENERATED SDK for me API

// Information about a Bill entry
type RefundDetail struct {
	Description    string `json:"description,omitempty"`
	Domain         string `json:"domain,omitempty"`
	Quantity       string `json:"quantity,omitempty"`
	Reference      string `json:"reference,omitempty"`
	RefundDetailId string `json:"refundDetailId,omitempty"`
	RefundId       string `json:"refundId,omitempty"`
	TotalPrice     Price  `json:"totalPrice,omitempty"`
	UnitPrice      Price  `json:"unitPrice,omitempty"`
}
