package cloud

// GENERATED SDK for cloud API

// MonthlyResources
type MonthlyResources struct {
	Certification  []MonthlyCertification  `json:"certification,omitempty"`
	Instance       []MonthlyInstance       `json:"instance,omitempty"`
	InstanceOption []MonthlyInstanceOption `json:"instanceOption,omitempty"`
}
