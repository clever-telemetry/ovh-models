package cloud

// GENERATED SDK for cloud API

// AI Solutions Resource Object
type Resources struct {
	Cpu              int64   `json:"cpu,omitempty"`
	EphemeralStorage int64   `json:"ephemeralStorage,omitempty"`
	Flavor           string  `json:"flavor,omitempty"`
	Gpu              int64   `json:"gpu,omitempty"`
	GpuBrand         *string `json:"gpuBrand,omitempty"`
	GpuMemory        *int64  `json:"gpuMemory,omitempty"`
	GpuModel         *string `json:"gpuModel,omitempty"`
	Memory           int64   `json:"memory,omitempty"`
	PrivateNetwork   int64   `json:"privateNetwork,omitempty"`
	PublicNetwork    int64   `json:"publicNetwork,omitempty"`
}
