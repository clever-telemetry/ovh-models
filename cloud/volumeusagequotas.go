package cloud

// GENERATED SDK for cloud API

// Quotas on volumes
type VolumeUsageQuotas struct {
	MaxGigabytes  int64 `json:"maxGigabytes,omitempty"`
	UsedGigabytes int64 `json:"usedGigabytes,omitempty"`
	VolumeCount   int64 `json:"volumeCount,omitempty"`
}
