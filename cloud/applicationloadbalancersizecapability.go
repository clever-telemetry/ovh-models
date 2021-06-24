package cloud

// GENERATED SDK for cloud API

// Application load balancer size capability
type ApplicationLoadBalancerSizeCapability struct {
	BandwidthMbPerSecond int64    `json:"bandwidthMbPerSecond,omitempty"`
	MaximumConnection    int64    `json:"maximumConnection,omitempty"`
	RequestsPerSecond    int64    `json:"requestsPerSecond,omitempty"`
	Size                 SizeEnum `json:"size,omitempty"`
}
