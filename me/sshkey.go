package me

// GENERATED SDK for me API

// Customer public SSH key, can be used for rescue netboot or server access after reinstallation
type sshKey struct {
	Default bool   `json:"default,omitempty"`
	Key     string `json:"key,omitempty"`
	KeyName string `json:"keyName,omitempty"`
}
