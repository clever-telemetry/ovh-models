package cloud

// GENERATED SDK for cloud API

// Network load balancer size capability
type LoadBalancerSizeCapability struct {
	BandwidthMbPerSecond   int64    `json:"bandwidthMbPerSecond,omitempty"`
	MaximumConnection      int64    `json:"maximumConnection,omitempty"`
	NewConnectionPerSecond int64    `json:"newConnectionPerSecond,omitempty"`
	Size                   SizeEnum `json:"size,omitempty"`
}
