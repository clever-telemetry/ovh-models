package me

// GENERATED SDK for me API

// Commercial offer a customer can migrate his service to
type ProposedOffer struct {
	Configurations []ProposedOfferConfiguration `json:"configurations,omitempty"`
	Plan           GenericProductDefinition     `json:"plan,omitempty"`
	PricingMode    string                       `json:"pricingMode,omitempty"`
	Promotion      PercentagePromotion          `json:"promotion,omitempty"`
}
