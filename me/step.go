package me

// GENERATED SDK for me API

// Country Migration Step
type Step struct {
	Contracts *Contracts `json:"contracts,omitempty"`
	Debt      *Debt      `json:"debt,omitempty"`
	Name      NameEnum   `json:"name,omitempty"`
	Orders    *Orders    `json:"orders,omitempty"`
	Status    StatusEnum `json:"status,omitempty"`
}
