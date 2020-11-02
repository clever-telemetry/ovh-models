package cloud

// GENERATED SDK for cloud API

// Volumes usage for current month
type VolumesUsage struct {
	Detail []VolumeUsageDetail `json:"detail,omitempty"`
	Total  Price               `json:"total,omitempty"`
}
