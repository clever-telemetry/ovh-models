package me

// GENERATED SDK for me API

// Details about an OVH account
type OvhAccount struct {
	AlertThreshold *int64 `json:"alertThreshold,omitempty"`
	Balance        Price  `json:"balance,omitempty"`
	CanBeCredited  bool   `json:"canBeCredited,omitempty"`
	IsActive       bool   `json:"isActive,omitempty"`
	LastUpdate     string `json:"lastUpdate,omitempty"`
	OpenDate       string `json:"openDate,omitempty"`
	OvhAccountId   string `json:"ovhAccountId,omitempty"`
}
