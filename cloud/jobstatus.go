package cloud

// GENERATED SDK for cloud API

// AI Solutions Job Status Object
type JobStatus struct {
	DataSync           []DataSync         `json:"dataSync,omitempty"`
	Duration           *int64             `json:"duration,omitempty"`
	ExitCode           *int64             `json:"exitCode,omitempty"`
	FinalizedAt        *string            `json:"finalizedAt,omitempty"`
	History            []JobStatusHistory `json:"history,omitempty"`
	Infos              *string            `json:"infos,omitempty"`
	InitializingAt     *string            `json:"initializingAt,omitempty"`
	Ip                 *string            `json:"ip,omitempty"`
	JobUrl             *string            `json:"jobUrl,omitempty"`
	LastTransitionDate *string            `json:"lastTransitionDate,omitempty"`
	MonitoringUrl      *string            `json:"monitoringUrl,omitempty"`
	QueuedAt           *string            `json:"queuedAt,omitempty"`
	SshUrl             *string            `json:"sshUrl,omitempty"`
	StartedAt          *string            `json:"startedAt,omitempty"`
	State              *JobStateEnum      `json:"state,omitempty"`
	StoppedAt          *string            `json:"stoppedAt,omitempty"`
	Url                *string            `json:"url,omitempty"`
}
