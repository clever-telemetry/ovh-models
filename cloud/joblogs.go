package cloud

// GENERATED SDK for cloud API

// Job Logs
type JobLogs struct {
	Logs        []LogLine `json:"logs,omitempty"`
	LogsAddress *string   `json:"logsAddress,omitempty"`
	StartDate   string    `json:"startDate,omitempty"`
}
