package cloud

// GENERATED SDK for cloud API

// Quotas for network
type NetworkQuota struct {
	Networks int64 `json:"networks,omitempty"`
	Ports    int64 `json:"ports,omitempty"`
	Subnets  int64 `json:"subnets,omitempty"`
}
