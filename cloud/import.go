package cloud

// GENERATED SDK for cloud API

// Import external certificate
type Import struct {
	Cert  string  `json:"cert"`
	Chain *string `json:"chain,omitempty"`
	Key   string  `json:"key"`
}
