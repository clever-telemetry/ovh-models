package cloud

// GENERATED SDK for cloud API

// Specific database engine capability
type Engine struct {
	DefaultVersion string   `json:"defaultVersion,omitempty"`
	Description    string   `json:"description,omitempty"`
	Name           string   `json:"name,omitempty"`
	Versions       []string `json:"versions,omitempty"`
}
