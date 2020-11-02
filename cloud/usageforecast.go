package cloud

// GENERATED SDK for cloud API

// UsageForecast
type UsageForecast struct {
	HourlyUsage    *HourlyResources  `json:"hourlyUsage,omitempty"`
	LastUpdate     string            `json:"lastUpdate,omitempty"`
	MonthlyUsage   *MonthlyResources `json:"monthlyUsage,omitempty"`
	Period         Period            `json:"period,omitempty"`
	ResourcesUsage *[]TypedResources `json:"resourcesUsage,omitempty"`
	UsableCredits  *UsedCredits      `json:"usableCredits,omitempty"`
}
