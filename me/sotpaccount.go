package me

// GENERATED SDK for me API

// SOTP Two-Factor Authentication
type SOTPAccount struct {
	CreationDate string         `json:"creationDate,omitempty"`
	LastUsedDate *string        `json:"lastUsedDate,omitempty"`
	Remaining    int64          `json:"remaining,omitempty"`
	Status       SOTPStatusEnum `json:"status,omitempty"`
}
