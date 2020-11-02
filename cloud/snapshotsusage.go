package cloud

// GENERATED SDK for cloud API

// Snapshots usage for current month
type SnapshotsUsage struct {
	Detail []SnapshotUsageDetail `json:"detail,omitempty"`
	Total  Price                 `json:"total,omitempty"`
}
