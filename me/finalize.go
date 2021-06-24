package me

// GENERATED SDK for me API

// Payload to finalize payment method registration
type Finalize struct {
	ExpirationMonth *int64  `json:"expirationMonth,omitempty"`
	ExpirationYear  *int64  `json:"expirationYear,omitempty"`
	FormSessionId   *string `json:"formSessionId,omitempty"`
	RegistrationId  *string `json:"registrationId,omitempty"`
}
