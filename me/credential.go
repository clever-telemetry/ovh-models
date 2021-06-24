package me

// GENERATED SDK for me API

// API Credential
type Credential struct {
	AllowedIPs    *[]string           `json:"allowedIPs,omitempty"`
	ApplicationId int64               `json:"applicationId,omitempty"`
	Creation      string              `json:"creation,omitempty"`
	CredentialId  int64               `json:"credentialId,omitempty"`
	Expiration    *string             `json:"expiration,omitempty"`
	LastUse       *string             `json:"lastUse,omitempty"`
	OvhSupport    bool                `json:"ovhSupport,omitempty"`
	Rules         []AccessRule        `json:"rules,omitempty"`
	Status        CredentialStateEnum `json:"status,omitempty"`
}
