package me

// GENERATED SDK for me API

// Price with its currency and textual representation
type Price struct {
	CurrencyCode CurrencyCodeEnum `json:"currencyCode,omitempty"`
	Text         string           `json:"text,omitempty"`
	Value        float64          `json:"value,omitempty"`
}
