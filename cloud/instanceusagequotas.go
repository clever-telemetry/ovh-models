package cloud

// GENERATED SDK for cloud API

// Quotas on instances
type InstanceUsageQuotas struct {
	MaxCores      int64 `json:"maxCores,omitempty"`
	MaxInstances  int64 `json:"maxInstances,omitempty"`
	MaxRam        int64 `json:"maxRam,omitempty"`
	UsedCores     int64 `json:"usedCores,omitempty"`
	UsedInstances int64 `json:"usedInstances,omitempty"`
	UsedRAM       int64 `json:"usedRAM,omitempty"`
}
