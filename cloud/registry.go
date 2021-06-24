package cloud

// GENERATED SDK for cloud API

// Representation of a registry
type Registry struct {
	Custom   bool    `json:"custom,omitempty"`
	Password *string `json:"password"`
	Url      *string `json:"url"`
	Username *string `json:"username"`
}
