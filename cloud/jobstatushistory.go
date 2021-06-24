package cloud

// GENERATED SDK for cloud API

// AI Solutions Job Status History Object
type JobStatusHistory struct {
	Date  string       `json:"date,omitempty"`
	State JobStateEnum `json:"state,omitempty"`
}
