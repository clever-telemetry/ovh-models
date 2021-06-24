package cloud

// GENERATED SDK for cloud API

// Cloud Storage Quota
type Quota struct {
	BytesUsed      int64  `json:"bytesUsed,omitempty"`
	ContainerCount int64  `json:"containerCount,omitempty"`
	ObjectCount    int64  `json:"objectCount,omitempty"`
	QuotaBytes     *int64 `json:"quotaBytes,omitempty"`
}
