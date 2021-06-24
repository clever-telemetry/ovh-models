package me

// GENERATED SDK for me API

// Plan data from order
type OrderPlan struct {
	Code        *string          `json:"code,omitempty"`
	Duration    *string          `json:"duration,omitempty"`
	PricingMode *string          `json:"pricingMode,omitempty"`
	Product     OrderPlanProduct `json:"product,omitempty"`
	Quantity    *int64           `json:"quantity,omitempty"`
}
