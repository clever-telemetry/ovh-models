package me

// GENERATED SDK for me API

// Element of consumption for resource
type Element struct {
	Details    []Detail `json:"details,omitempty"`
	PlanCode   string   `json:"planCode,omitempty"`
	PlanFamily string   `json:"planFamily,omitempty"`
	Price      Price    `json:"price,omitempty"`
	Quantity   int64    `json:"quantity,omitempty"`
}
