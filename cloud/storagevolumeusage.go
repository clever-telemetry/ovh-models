package cloud

// GENERATED SDK for cloud API

// Storage volume used on your project
type StorageVolumeUsage struct {
	Region      string `json:"region,omitempty"`
	StoredBytes int64  `json:"storedBytes,omitempty"`
	Total       Price  `json:"total,omitempty"`
}
