package cloud

// GENERATED SDK for cloud API

// AI Solutions Global Resource per flavor unit
type ResourcesPerUnit struct {
	Cpu              int64 `json:"cpu,omitempty"`
	EphemeralStorage int64 `json:"ephemeralStorage,omitempty"`
	Memory           int64 `json:"memory,omitempty"`
	PrivateNetwork   int64 `json:"privateNetwork,omitempty"`
	PublicNetwork    int64 `json:"publicNetwork,omitempty"`
}
