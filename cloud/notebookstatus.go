package cloud

// GENERATED SDK for cloud API

// AI Solutions Notebook Status Object
type NotebookStatus struct {
	DataSync      []DataSync         `json:"dataSync,omitempty"`
	Duration      *int64             `json:"duration,omitempty"`
	Infos         *string            `json:"infos,omitempty"`
	LastStartedAt *string            `json:"lastStartedAt,omitempty"`
	LastStoppedAt *string            `json:"lastStoppedAt,omitempty"`
	MonitoringUrl *string            `json:"monitoringUrl,omitempty"`
	State         *NotebookStateEnum `json:"state,omitempty"`
	Url           *string            `json:"url,omitempty"`
}
