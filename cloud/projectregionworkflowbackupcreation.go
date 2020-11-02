package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectRegionWorkflowBackupCreation struct {
	Cron              string `json:"cron"`
	InstanceId        string `json:"instanceId"`
	MaxExecutionCount *int64 `json:"maxExecutionCount,omitempty"`
	Name              string `json:"name"`
	Rotation          int64  `json:"rotation"`
}
