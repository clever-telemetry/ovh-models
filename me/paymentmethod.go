package me

// GENERATED SDK for me API

// Payment method object
type PaymentMethod struct {
	BillingContactId *int64     `json:"billingContactId,omitempty"`
	CreationDate     string     `json:"creationDate,omitempty"`
	Default          bool       `json:"default,omitempty"`
	Description      *string    `json:"description,omitempty"`
	ExpirationDate   *string    `json:"expirationDate,omitempty"`
	Icon             Icon       `json:"icon,omitempty"`
	Label            *string    `json:"label,omitempty"`
	LastUpdate       string     `json:"lastUpdate,omitempty"`
	PaymentMeanId    *int64     `json:"paymentMeanId,omitempty"`
	PaymentMethodId  int64      `json:"paymentMethodId,omitempty"`
	PaymentSubType   *string    `json:"paymentSubType,omitempty"`
	PaymentType      string     `json:"paymentType,omitempty"`
	Status           StatusEnum `json:"status,omitempty"`
}
