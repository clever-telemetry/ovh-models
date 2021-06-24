package me

// GENERATED SDK for me API

// Missing description
type ExpiringMovement struct {
	Amount         Price     `json:"amount,omitempty"`
	CreationDate   string    `json:"creationDate,omitempty"`
	ExpirationDate string    `json:"expirationDate,omitempty"`
	LastUpdate     string    `json:"lastUpdate,omitempty"`
	SourceObject   SubObject `json:"sourceObject,omitempty"`
}
