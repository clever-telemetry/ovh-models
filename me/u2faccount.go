package me

// GENERATED SDK for me API

// U2F Two-Factor Authentication
type U2FAccount struct {
	CreationDate string        `json:"creationDate,omitempty"`
	Description  string        `json:"description,omitempty"`
	Id           int64         `json:"id,omitempty"`
	LastUsedDate *string       `json:"lastUsedDate,omitempty"`
	Status       U2FStatusEnum `json:"status,omitempty"`
}
