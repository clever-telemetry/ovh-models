package me

// GENERATED SDK for me API

// Describe TOTP secret keys
type TOTPSecret struct {
	Id           int64  `json:"id,omitempty"`
	QrcodeHelper string `json:"qrcodeHelper,omitempty"`
	Secret       string `json:"secret,omitempty"`
}
