package cloud

// GENERATED SDK for cloud API

// An operation is an async process on your Project
type Operation struct {
	Action      string              `json:"action,omitempty"`
	CompletedAt *string             `json:"completedAt,omitempty"`
	CreatedAt   string              `json:"createdAt,omitempty"`
	Id          string              `json:"id,omitempty"`
	Progress    int64               `json:"progress,omitempty"`
	Regions     *[]string           `json:"regions,omitempty"`
	StartedAt   *string             `json:"startedAt,omitempty"`
	Status      OperationStatusEnum `json:"status,omitempty"`
}
