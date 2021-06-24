package me

// GENERATED SDK for me API

// Registration response to validate
type Validation struct {
	FormSessionId   *string         `json:"formSessionId,omitempty"`
	MerchandId      *string         `json:"merchandId,omitempty"`
	OrganizationId  *string         `json:"organizationId,omitempty"`
	PaymentMethodId int64           `json:"paymentMethodId,omitempty"`
	TransactionId   *int64          `json:"transactionId,omitempty"`
	Url             *string         `json:"url,omitempty"`
	ValidationType  IntegrationEnum `json:"validationType,omitempty"`
}
