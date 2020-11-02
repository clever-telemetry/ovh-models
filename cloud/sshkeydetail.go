package cloud

// GENERATED SDK for cloud API

// SshKeyDetail
type SshKeyDetail struct {
	FingerPrint string   `json:"fingerPrint,omitempty"`
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	PublicKey   string   `json:"publicKey,omitempty"`
	Regions     []string `json:"regions,omitempty"`
}
