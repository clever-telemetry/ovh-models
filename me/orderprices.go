package me

// GENERATED SDK for me API

// Prices of an order
type OrderPrices struct {
	OriginalWithoutTax *Price `json:"originalWithoutTax,omitempty"`
	Reduction          *Price `json:"reduction,omitempty"`
	Tax                Price  `json:"tax,omitempty"`
	WithTax            Price  `json:"withTax,omitempty"`
	WithoutTax         Price  `json:"withoutTax,omitempty"`
}
