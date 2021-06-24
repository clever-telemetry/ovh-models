package me

// GENERATED SDK for me API

// Tag
type Tag struct {
	CreationDate   string     `json:"creationDate,omitempty"`
	CustomerCode   string     `json:"customerCode,omitempty"`
	ExpirationDate *string    `json:"expirationDate,omitempty"`
	Extra          *TagExtra  `json:"extra,omitempty"`
	LastUpdate     string     `json:"lastUpdate,omitempty"`
	Status         StatusEnum `json:"status,omitempty"`
	Tag            string     `json:"tag,omitempty"`
}
