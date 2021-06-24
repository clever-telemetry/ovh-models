package me

// GENERATED SDK for me API

// Information about a Bill entry
type BillDetail struct {
	BillDetailId string  `json:"billDetailId,omitempty"`
	Description  string  `json:"description,omitempty"`
	Domain       string  `json:"domain,omitempty"`
	PeriodEnd    *string `json:"periodEnd,omitempty"`
	PeriodStart  *string `json:"periodStart,omitempty"`
	Quantity     string  `json:"quantity,omitempty"`
	TotalPrice   Price   `json:"totalPrice,omitempty"`
	UnitPrice    Price   `json:"unitPrice,omitempty"`
}
