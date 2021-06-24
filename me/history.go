package me

// GENERATED SDK for me API

// Step history of order follow-up
type History struct {
	Date        string            `json:"date,omitempty"`
	Description string            `json:"description,omitempty"`
	Label       HistoryStatusEnum `json:"label,omitempty"`
}
