package me

// GENERATED SDK for me API

// Paypal informations
type Paypal struct {
	AgreementId        string          `json:"agreementId,omitempty"`
	CreationDate       string          `json:"creationDate,omitempty"`
	DefaultPaymentMean bool            `json:"defaultPaymentMean,omitempty"`
	Description        *string         `json:"description,omitempty"`
	Email              string          `json:"email,omitempty"`
	Icon               *IconData       `json:"icon,omitempty"`
	Id                 int64           `json:"id,omitempty"`
	State              PaypalStateEnum `json:"state,omitempty"`
}
