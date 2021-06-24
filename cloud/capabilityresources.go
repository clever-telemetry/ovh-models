package cloud

// GENERATED SDK for cloud API

// AI Solutions Platform Resource Object
type CapabilityResources struct {
	DataRegions []string `json:"dataRegions,omitempty"`
	Gpus        []Gpu    `json:"gpus,omitempty"`
	MaxCpus     int64    `json:"maxCpus,omitempty"`
}
