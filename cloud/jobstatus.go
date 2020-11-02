package cloud

// GENERATED SDK for cloud API

// Training Platform Job Status Object
type JobStatus struct {
	Duration      *int64        `json:"duration,omitempty"`
	History       []JobStatus   `json:"history,omitempty"`
	Infos         *string       `json:"infos,omitempty"`
	Ip            *string       `json:"ip,omitempty"`
	JobUrl        *string       `json:"jobUrl,omitempty"`
	MonitoringUrl *string       `json:"monitoringUrl,omitempty"`
	QueuedAt      *string       `json:"queuedAt,omitempty"`
	StartedAt     *string       `json:"startedAt,omitempty"`
	State         *JobStateEnum `json:"state,omitempty"`
	StoppedAt     *string       `json:"stoppedAt,omitempty"`
}
