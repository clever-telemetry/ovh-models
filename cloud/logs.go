package cloud

// GENERATED SDK for cloud API

// Instance Logs
type Logs struct {
	LastActivity *string   `json:"lastActivity,omitempty"`
	Logs         []LogLine `json:"logs,omitempty"`
}
