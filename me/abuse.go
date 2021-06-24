package me

// GENERATED SDK for me API

// Get report API response
type Abuse struct {
	Category     AbuseCategoryEnum `json:"category,omitempty"`
	CreationDate string            `json:"creationDate,omitempty"`
	PublicId     string            `json:"publicId,omitempty"`
	Service      string            `json:"service,omitempty"`
	Status       AbuseStatusEnum   `json:"status,omitempty"`
}
