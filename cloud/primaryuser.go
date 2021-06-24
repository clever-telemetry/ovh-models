package cloud

// GENERATED SDK for cloud API

// Mongodb primary user definition
type PrimaryUser struct {
	Password string   `json:"password,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Username string   `json:"username,omitempty"`
}
