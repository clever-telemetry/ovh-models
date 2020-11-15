package auth

// GENERATED SDK for auth API

// Details about the authentication used
type Details struct {
	Description *string    `json:"description,omitempty"`
	Method      MethodEnum `json:"method,omitempty"`
	Roles       *[]string  `json:"roles,omitempty"`
	User        *string    `json:"user,omitempty"`
}
