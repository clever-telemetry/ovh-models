package cloud

// GENERATED SDK for cloud API

// UsageCurrent
type UsageCurrent struct {
	HourlyUsage    *HourlyResources  `json:"hourlyUsage,omitempty"`
	LastUpdate     string            `json:"lastUpdate,omitempty"`
	MonthlyUsage   *MonthlyResources `json:"monthlyUsage,omitempty"`
	Period         Period            `json:"period,omitempty"`
	ResourcesUsage *[]TypedResources `json:"resourcesUsage,omitempty"`
}
