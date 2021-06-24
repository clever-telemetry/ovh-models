package me

// GENERATED SDK for me API

// The object linked to this debt entry
type AssociatedObject struct {
	Id          *string      `json:"id,omitempty"`
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty"`
	SubId       *string      `json:"subId,omitempty"`
	Type        *TypeEnum    `json:"type,omitempty"`
}
