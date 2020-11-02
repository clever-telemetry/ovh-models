package cloud

// GENERATED SDK for cloud API

// Quotas for compute
type ComputeQuota struct {
	Cores     int64 `json:"cores,omitempty"`
	Instances int64 `json:"instances,omitempty"`
	Ram       int64 `json:"ram,omitempty"`
}
