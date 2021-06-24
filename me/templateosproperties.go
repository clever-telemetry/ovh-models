package me

// GENERATED SDK for me API

// A structure describing properties customizables about this dedicated installation template
type TemplateOsProperties struct {
	CustomHostname               *string `json:"customHostname,omitempty"`
	PostInstallationScriptLink   *string `json:"postInstallationScriptLink,omitempty"`
	PostInstallationScriptReturn *string `json:"postInstallationScriptReturn,omitempty"`
	SshKeyName                   *string `json:"sshKeyName,omitempty"`
	UseDistributionKernel        *bool   `json:"useDistributionKernel,omitempty"`
}
