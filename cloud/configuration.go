package cloud

// GENERATED SDK for cloud API

// A load balancer configuration
type Configuration struct {
	Backends     []Backend  `json:"backends,omitempty"`
	Certificates []string   `json:"certificates,omitempty"`
	Frontends    []Frontend `json:"frontends,omitempty"`
	Networking   Networking `json:"networking,omitempty"`
	Version      int64      `json:"version,omitempty"`
}
