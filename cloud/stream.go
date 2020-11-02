package cloud

// GENERATED SDK for cloud API

// A stream to send data
type Stream struct {
	Backlog     string           `json:"backlog,omitempty"`
	Description *string          `json:"description,omitempty"`
	Id          string           `json:"id,omitempty"`
	Kind        StreamKindEnum   `json:"kind,omitempty"`
	Name        string           `json:"name,omitempty"`
	Regions     []string         `json:"regions,omitempty"`
	Retention   string           `json:"retention,omitempty"`
	Status      StreamStatusEnum `json:"status,omitempty"`
	Throttling  int64            `json:"throttling,omitempty"`
}
