package cloud

// GENERATED SDK for cloud API

// Cloud alerting consumption
type Alerting struct {
	CreationDate              string            `json:"creationDate,omitempty"`
	Delay                     AlertingDelayEnum `json:"delay,omitempty"`
	Email                     string            `json:"email,omitempty"`
	FormattedMonthlyThreshold Price             `json:"formattedMonthlyThreshold,omitempty"`
	Id                        string            `json:"id,omitempty"`
	MonthlyThreshold          int64             `json:"monthlyThreshold,omitempty"`
}
