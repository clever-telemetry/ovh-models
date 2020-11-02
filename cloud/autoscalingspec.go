package cloud

// GENERATED SDK for cloud API

// Autoscaling specification
type AutoscalingSpec struct {
	CpuAverageUtilization    *int64 `json:"cpuAverageUtilization,omitempty"`
	MaxReplicas              *int64 `json:"maxReplicas,omitempty"`
	MemoryAverageUtilization *int64 `json:"memoryAverageUtilization,omitempty"`
	MinReplicas              *int64 `json:"minReplicas,omitempty"`
}
