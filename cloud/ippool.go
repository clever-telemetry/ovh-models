package cloud

// GENERATED SDK for cloud API

// IPPool
type IPPool struct {
	Dhcp    bool   `json:"dhcp,omitempty"`
	End     string `json:"end,omitempty"`
	Network string `json:"network,omitempty"`
	Region  string `json:"region,omitempty"`
	Start   string `json:"start,omitempty"`
}
