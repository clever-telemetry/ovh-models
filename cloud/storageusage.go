package cloud

// GENERATED SDK for cloud API

// Usage information for current month on your project
type StorageUsage struct {
	Bandwidth []BandwidthStorageUsage `json:"bandwidth,omitempty"`
	Total     Price                   `json:"total,omitempty"`
	Volume    []StorageVolumeUsage    `json:"volume,omitempty"`
}
