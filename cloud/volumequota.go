package cloud

// GENERATED SDK for cloud API

// Quotas for volume
type VolumeQuota struct {
	Gigabytes int64 `json:"gigabytes,omitempty"`
	Snapshots int64 `json:"snapshots,omitempty"`
	Volumes   int64 `json:"volumes,omitempty"`
}
