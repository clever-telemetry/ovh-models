package me

// GENERATED SDK for me API

// Register validation payload result
type ValidationResult struct {
	FormSessionId   *string         `json:"formSessionId,omitempty"`
	MerchantId      *string         `json:"merchantId,omitempty"`
	OrganizationId  *string         `json:"organizationId,omitempty"`
	PaymentMethodId int64           `json:"paymentMethodId,omitempty"`
	TransactionId   int64           `json:"transactionId,omitempty"`
	Url             *string         `json:"url,omitempty"`
	ValidationType  IntegrationType `json:"validationType,omitempty"`
}
