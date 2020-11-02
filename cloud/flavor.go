package cloud

// GENERATED SDK for cloud API

// Compute Flavor for the Serving Engine
type Flavor struct {
	CpuCore     int64  `json:"cpuCore,omitempty"`
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	RamMB       int64  `json:"ramMB,omitempty"`
}
