package me

// GENERATED SDK for me API

// Follow up history of an order
type FollowUp struct {
	History []History  `json:"history,omitempty"`
	Status  StatusEnum `json:"status,omitempty"`
	Step    StepEnum   `json:"step,omitempty"`
}
