package me

// GENERATED SDK for me API

// Representation of country and continent from visitor IP
type ContinentCountryLocation struct {
	Continent   ContinentEnum `json:"continent,omitempty"`
	CountryCode CountryEnum   `json:"countryCode,omitempty"`
	Ip          string        `json:"ip,omitempty"`
}
