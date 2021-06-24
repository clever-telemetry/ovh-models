package me

// GENERATED SDK for me API

// A user
type User struct {
	Creation           string     `json:"creation,omitempty"`
	Description        string     `json:"description,omitempty"`
	Email              string     `json:"email,omitempty"`
	Group              string     `json:"group,omitempty"`
	LastUpdate         string     `json:"lastUpdate,omitempty"`
	Login              string     `json:"login,omitempty"`
	PasswordLastUpdate string     `json:"passwordLastUpdate,omitempty"`
	Status             UserStatus `json:"status,omitempty"`
}
