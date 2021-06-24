package me

// GENERATED SDK for me API

// TOTP Two-Factor Authentication
type TOTPAccount struct {
	CreationDate string         `json:"creationDate,omitempty"`
	Description  string         `json:"description,omitempty"`
	Id           int64          `json:"id,omitempty"`
	LastUsedDate *string        `json:"lastUsedDate,omitempty"`
	Status       TOTPStatusEnum `json:"status,omitempty"`
}
