package cloud

// GENERATED SDK for cloud API

// Usage information for current month on your project
type BandwidthStorageUsage struct {
	DownloadedBytes int64  `json:"downloadedBytes,omitempty"`
	Region          string `json:"region,omitempty"`
	Total           Price  `json:"total,omitempty"`
}
