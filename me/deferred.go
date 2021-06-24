package me

// GENERATED SDK for me API

// Deferred account informations
type Deferred struct {
	CreationDate       string            `json:"creationDate,omitempty"`
	DefaultPaymentMean bool              `json:"defaultPaymentMean,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Icon               *IconData         `json:"icon,omitempty"`
	Id                 int64             `json:"id,omitempty"`
	Label              *string           `json:"label,omitempty"`
	State              DeferredStateEnum `json:"state,omitempty"`
}
