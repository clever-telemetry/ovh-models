package cloud

// GENERATED SDK for cloud API

// Quotas
type Quotas struct {
	Instance *InstanceUsageQuotas `json:"instance,omitempty"`
	Keypair  *KeypairQuotas       `json:"keypair,omitempty"`
	Region   string               `json:"region,omitempty"`
	Volume   *VolumeUsageQuotas   `json:"volume,omitempty"`
}
