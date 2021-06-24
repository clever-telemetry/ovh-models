package me

// GENERATED SDK for me API

// Order detail reduction
type Reduction struct {
	Context ReductionContextEnum `json:"context,omitempty"`
	Price   Price                `json:"price,omitempty"`
	Type    ReductionTypeEnum    `json:"type,omitempty"`
	Value   Price                `json:"value,omitempty"`
}
