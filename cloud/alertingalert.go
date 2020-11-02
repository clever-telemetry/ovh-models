package cloud

// GENERATED SDK for cloud API

// Cloud alert on your consumption
type AlertingAlert struct {
	AlertDate string   `json:"alertDate,omitempty"`
	AlertId   int64    `json:"alertId,omitempty"`
	Emails    []string `json:"emails,omitempty"`
}
