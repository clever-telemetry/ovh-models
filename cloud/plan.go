package cloud

// GENERATED SDK for cloud API

// Plan of the registry
type Plan struct {
	Code           string   `json:"code,omitempty"`
	CreatedAt      string   `json:"createdAt,omitempty"`
	Features       Features `json:"features,omitempty"`
	Id             string   `json:"id,omitempty"`
	Name           string   `json:"name,omitempty"`
	RegistryLimits Limits   `json:"registryLimits,omitempty"`
	UpdatedAt      string   `json:"updatedAt,omitempty"`
}
