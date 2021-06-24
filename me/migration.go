package me

// GENERATED SDK for me API

// Country Migration
type Migration struct {
	From   OvhSubsidiaryEnum `json:"from,omitempty"`
	Id     int64             `json:"id,omitempty"`
	Status StatusEnum        `json:"status,omitempty"`
	Steps  *[]Step           `json:"steps,omitempty"`
	To     OvhSubsidiaryEnum `json:"to,omitempty"`
}
