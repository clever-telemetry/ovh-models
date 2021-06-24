package me

// GENERATED SDK for me API

// Tag configuration key
type TagKey struct {
	Enum      *[]string `json:"enum,omitempty"`
	Key       string    `json:"key,omitempty"`
	MaxLength *int64    `json:"maxLength,omitempty"`
	MinValue  *int64    `json:"minValue,omitempty"`
	Optional  bool      `json:"optional,omitempty"`
	Pattern   *string   `json:"pattern,omitempty"`
	Type      TypeEnum  `json:"type,omitempty"`
}
