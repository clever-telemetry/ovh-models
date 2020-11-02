package cloud

// GENERATED SDK for cloud API

// An execution of the backup workflow
type Execution struct {
	ExecutedAt string             `json:"executedAt,omitempty"`
	State      ExecutionStateEnum `json:"state,omitempty"`
	StateInfo  string             `json:"stateInfo,omitempty"`
}
