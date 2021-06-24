package me

// GENERATED SDK for me API

// An order
type Order struct {
	Contracts []Contract    `json:"contracts,omitempty"`
	Details   []OrderDetail `json:"details,omitempty"`
	OrderId   *int64        `json:"orderId,omitempty"`
	Prices    OrderPrices   `json:"prices,omitempty"`
	Url       *string       `json:"url,omitempty"`
}
