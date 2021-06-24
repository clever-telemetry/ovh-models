package cloud

// GENERATED SDK for cloud API

// A load balancer entryPoint
type EntryPoint struct {
	DefaultTarget *string      `json:"defaultTarget,omitempty"`
	DisableH2     *bool        `json:"disableH2,omitempty"`
	Name          string       `json:"name"`
	PortRanges    *[]PortRange `json:"portRanges,omitempty"`
	Ports         *[]int64     `json:"ports,omitempty"`
	Rules         []Rule       `json:"rules,omitempty"`
	Tls           bool         `json:"tls,omitempty"`
}
