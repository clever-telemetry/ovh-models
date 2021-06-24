package me

// GENERATED SDK for me API

// Details about an IP block organisation
type Ipv4Org struct {
	Abuse_mailbox  string         `json:"abuse_mailbox,omitempty"`
	Address        string         `json:"address,omitempty"`
	City           string         `json:"city,omitempty"`
	Country        CountryEnum    `json:"country,omitempty"`
	Firstname      string         `json:"firstname,omitempty"`
	Lastname       string         `json:"lastname,omitempty"`
	OrganisationId string         `json:"organisationId,omitempty"`
	Phone          interface{}    `json:"phone,omitempty"`
	Registry       IpRegistryEnum `json:"registry,omitempty"`
	State          *string        `json:"state,omitempty"`
	Zip            *string        `json:"zip,omitempty"`
}
