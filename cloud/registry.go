package cloud

// GENERATED SDK for cloud API

// Representation of a registry
type Registry struct {
	CreatedAt string  `json:"createdAt,omitempty"`
	Id        string  `json:"id,omitempty"`
	Password  *string `json:"password"`
	Region    string  `json:"region"`
	UpdatedAt string  `json:"updatedAt,omitempty"`
	Url       *string `json:"url"`
	User      string  `json:"user,omitempty"`
	Username  *string `json:"username"`
}
