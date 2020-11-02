package cloud

// GENERATED SDK for cloud API

// Engine parameters
type CapabilitiesEngineParameter struct {
	Default     *string            `json:"default,omitempty"`
	Description string             `json:"description,omitempty"`
	Mandatory   bool               `json:"mandatory,omitempty"`
	Name        string             `json:"name,omitempty"`
	Type        string             `json:"type,omitempty"`
	Validator   ParameterValidator `json:"validator,omitempty"`
}
