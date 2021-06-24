package me

// GENERATED SDK for me API

// Representation of a product pricing
type GenericProductPricing struct {
	Capacities      []GenericProductPricingCapacitiesEnum `json:"capacities,omitempty"`
	Description     string                                `json:"description,omitempty"`
	Duration        string                                `json:"duration,omitempty"`
	Interval        int64                                 `json:"interval,omitempty"`
	MaximumQuantity *int64                                `json:"maximumQuantity,omitempty"`
	MaximumRepeat   *int64                                `json:"maximumRepeat,omitempty"`
	MinimumQuantity int64                                 `json:"minimumQuantity,omitempty"`
	MinimumRepeat   int64                                 `json:"minimumRepeat,omitempty"`
	Price           Price                                 `json:"price,omitempty"`
	PriceInUcents   int64                                 `json:"priceInUcents,omitempty"`
	PricingMode     string                                `json:"pricingMode,omitempty"`
	PricingType     GenericProductPricingTypeEnum         `json:"pricingType,omitempty"`
}
