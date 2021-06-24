package me

// GENERATED SDK for me API

// Representation of an Address
type Address struct {
	City         string      `json:"city,omitempty"`
	Country      CountryEnum `json:"country,omitempty"`
	Line1        string      `json:"line1,omitempty"`
	Line2        *string     `json:"line2,omitempty"`
	Line3        *string     `json:"line3,omitempty"`
	OtherDetails *string     `json:"otherDetails,omitempty"`
	Province     *string     `json:"province,omitempty"`
	Zip          string      `json:"zip,omitempty"`
}
