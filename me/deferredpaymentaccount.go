package me

// GENERATED SDK for me API

// Deferred payment account info
type DeferredPaymentAccount struct {
	CreationDate       string                           `json:"creationDate,omitempty"`
	DefaultPaymentMean bool                             `json:"defaultPaymentMean,omitempty"`
	Description        *string                          `json:"description,omitempty"`
	Icon               *IconData                        `json:"icon,omitempty"`
	Id                 int64                            `json:"id,omitempty"`
	Label              *string                          `json:"label,omitempty"`
	State              DeferredPaymentAccountStatusEnum `json:"state,omitempty"`
}
