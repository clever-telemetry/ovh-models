package me

// GENERATED SDK for me API

// The payment infos linked to this debt entry
type PaymentInfo struct {
	Description *string         `json:"description,omitempty"`
	PaymentType PaymentMeanEnum `json:"paymentType,omitempty"`
	PublicLabel *string         `json:"publicLabel,omitempty"`
}
