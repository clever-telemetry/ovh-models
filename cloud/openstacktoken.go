package cloud

// GENERATED SDK for cloud API

// OpenstackToken
type OpenstackToken struct {
	Catalog    []Catalog    `json:"catalog,omitempty"`
	Expires_at string       `json:"expires_at,omitempty"`
	Issued_at  string       `json:"issued_at,omitempty"`
	Methods    []string     `json:"methods,omitempty"`
	Project    TokenProject `json:"project,omitempty"`
	Roles      []Role       `json:"roles,omitempty"`
	User       UserToken    `json:"user,omitempty"`
}
