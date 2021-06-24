package cloud

// GENERATED SDK for cloud API

// A load balancer frontend
type Frontend struct {
	Backends   []BackendSelector `json:"backends"`
	Mode       *ModeEnum         `json:"mode,omitempty"`
	Name       string            `json:"name"`
	Port       *int64            `json:"port,omitempty"`
	PortRanges *[]PortRange      `json:"portRanges,omitempty"`
	Ports      *[]int64          `json:"ports,omitempty"`
	Tls        bool              `json:"tls,omitempty"`
	Whitelist  []string          `json:"whitelist,omitempty"`
}
