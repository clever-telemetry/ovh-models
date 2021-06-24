package cloud

// GENERATED SDK for cloud API

// Mongodb user definition
type UserCreation struct {
	Name     string   `json:"name,omitempty"`
	Password string   `json:"password,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}
