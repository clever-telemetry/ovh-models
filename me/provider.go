package me

// GENERATED SDK for me API

// A SAML 2.0 provider
type Provider struct {
	Creation              string      `json:"creation,omitempty"`
	GroupAttributeName    string      `json:"groupAttributeName,omitempty"`
	IdpSigningCertificate Certificate `json:"idpSigningCertificate,omitempty"`
	LastUpdate            string      `json:"lastUpdate,omitempty"`
	SsoServiceUrl         string      `json:"ssoServiceUrl,omitempty"`
}
