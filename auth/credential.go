package auth

// GENERATED SDK for auth API

// Credential request to get access to the API
type Credential struct {
	ConsumerKey   string              `json:"consumerKey,omitempty"`
	State         CredentialStateEnum `json:"state,omitempty"`
	ValidationUrl *string             `json:"validationUrl,omitempty"`
}
