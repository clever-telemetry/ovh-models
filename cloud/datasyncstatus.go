package cloud

// GENERATED SDK for cloud API

// AI Solutions Data Sync Status
type DataSyncStatus struct {
	EndedAt   *string           `json:"endedAt,omitempty"`
	Infos     *string           `json:"infos,omitempty"`
	Progress  []Progress        `json:"progress,omitempty"`
	QueuedAt  string            `json:"queuedAt,omitempty"`
	StartedAt *string           `json:"startedAt,omitempty"`
	State     DataSyncStateEnum `json:"state,omitempty"`
}
