package me

// GENERATED SDK for me API

// Balance of the fidelity account
type FidelityAccount struct {
	AlertThreshold *int64 `json:"alertThreshold,omitempty"`
	Balance        int64  `json:"balance,omitempty"`
	CanBeCredited  bool   `json:"canBeCredited,omitempty"`
	LastUpdate     string `json:"lastUpdate,omitempty"`
	OpenDate       string `json:"openDate,omitempty"`
}
