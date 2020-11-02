package cloud

// GENERATED SDK for cloud API

// SshKey
type SshKey struct {
	Id        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	PublicKey string   `json:"publicKey,omitempty"`
	Regions   []string `json:"regions,omitempty"`
}
