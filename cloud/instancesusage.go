package cloud

// GENERATED SDK for cloud API

// Instances usage for current month
type InstancesUsage struct {
	Detail []InstanceUsageDetail `json:"detail,omitempty"`
	Total  Price                 `json:"total,omitempty"`
}
