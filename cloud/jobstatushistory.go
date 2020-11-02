package cloud

// GENERATED SDK for cloud API

// Training Platform Job Status History Object
type JobStatusHistory struct {
	Date  string       `json:"date,omitempty"`
	State JobStateEnum `json:"state,omitempty"`
}
