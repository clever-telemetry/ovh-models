package me

// GENERATED SDK for me API

// A group
type Group struct {
	Creation     string   `json:"creation,omitempty"`
	DefaultGroup bool     `json:"defaultGroup,omitempty"`
	Description  string   `json:"description,omitempty"`
	LastUpdate   string   `json:"lastUpdate,omitempty"`
	Name         string   `json:"name,omitempty"`
	Role         RoleEnum `json:"role,omitempty"`
}
