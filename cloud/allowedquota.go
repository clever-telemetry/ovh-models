package cloud

// GENERATED SDK for cloud API

// Quotas
type AllowedQuota struct {
	Compute ComputeQuota `json:"compute,omitempty"`
	Name    string       `json:"name,omitempty"`
	Network NetworkQuota `json:"network,omitempty"`
	Volume  VolumeQuota  `json:"volume,omitempty"`
}
