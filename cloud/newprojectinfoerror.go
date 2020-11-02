package cloud

// GENERATED SDK for cloud API

// Error that can occur when creating a Public Cloud project
type NewProjectInfoError struct {
	Code    NewProjectInfoErrorCodeEnum `json:"code,omitempty"`
	Message string                      `json:"message,omitempty"`
}
