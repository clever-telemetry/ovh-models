package me

// GENERATED SDK for me API

// Detail of an order
type OrderDetail struct {
	CartItemID          *int64               `json:"cartItemID,omitempty"`
	Description         string               `json:"description,omitempty"`
	DetailType          *OrderDetailTypeEnum `json:"detailType,omitempty"`
	Domain              string               `json:"domain,omitempty"`
	OriginalTotalPrice  Price                `json:"originalTotalPrice,omitempty"`
	Quantity            int64                `json:"quantity,omitempty"`
	ReductionTotalPrice Price                `json:"reductionTotalPrice,omitempty"`
	Reductions          []Reduction          `json:"reductions,omitempty"`
	TotalPrice          Price                `json:"totalPrice,omitempty"`
	UnitPrice           Price                `json:"unitPrice,omitempty"`
}
