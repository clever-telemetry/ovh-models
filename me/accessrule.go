package me

// GENERATED SDK for me API

// Access rule required for the application
type AccessRule struct {
	Method MethodEnum `json:"method,omitempty"`
	Path   string     `json:"path,omitempty"`
}
