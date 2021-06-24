package me

// GENERATED SDK for me API

// Auto renewal information
type NicAutorenewInfos struct {
	Active    bool    `json:"active,omitempty"`
	LastRenew *string `json:"lastRenew,omitempty"`
	RenewDay  int64   `json:"renewDay,omitempty"`
}
