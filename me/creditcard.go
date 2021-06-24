package me

// GENERATED SDK for me API

// Credit card informations
type CreditCard struct {
	DefaultPaymentMean bool                `json:"defaultPaymentMean,omitempty"`
	Description        *string             `json:"description,omitempty"`
	ExpirationDate     string              `json:"expirationDate,omitempty"`
	Icon               *IconData           `json:"icon,omitempty"`
	Id                 int64               `json:"id,omitempty"`
	Number             string              `json:"number,omitempty"`
	State              CreditCardStateEnum `json:"state,omitempty"`
	ThreeDsValidated   bool                `json:"threeDsValidated,omitempty"`
	Type               string              `json:"type,omitempty"`
}
