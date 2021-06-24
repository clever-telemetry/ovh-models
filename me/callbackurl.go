package me

// GENERATED SDK for me API

// Callback URL's to register a new payment method
type CallbackUrl struct {
	Cancel  string `json:"cancel"`
	Error   string `json:"error"`
	Failure string `json:"failure"`
	Pending string `json:"pending"`
	Success string `json:"success"`
}
