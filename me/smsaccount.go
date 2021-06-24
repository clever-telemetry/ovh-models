package me

// GENERATED SDK for me API

// Sms Two-Factor Authentication
type SmsAccount struct {
	CreationDate string        `json:"creationDate,omitempty"`
	Description  string        `json:"description,omitempty"`
	Id           int64         `json:"id,omitempty"`
	LastUsedDate *string       `json:"lastUsedDate,omitempty"`
	PhoneNumber  string        `json:"phoneNumber,omitempty"`
	Status       SmsStatusEnum `json:"status,omitempty"`
}
