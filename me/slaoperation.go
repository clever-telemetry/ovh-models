package me

// GENERATED SDK for me API

// SLA properties
type SlaOperation struct {
	Date        string  `json:"date,omitempty"`
	Description string  `json:"description,omitempty"`
	EndDate     *string `json:"endDate,omitempty"`
	Id          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	StartDate   string  `json:"startDate,omitempty"`
}
