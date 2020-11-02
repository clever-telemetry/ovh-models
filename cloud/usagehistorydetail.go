package cloud

// GENERATED SDK for cloud API

// UsageHistoryDetail
type UsageHistoryDetail struct {
	HourlyUsage    *HourlyResources  `json:"hourlyUsage,omitempty"`
	Id             string            `json:"id,omitempty"`
	LastUpdate     string            `json:"lastUpdate,omitempty"`
	MonthlyUsage   *MonthlyResources `json:"monthlyUsage,omitempty"`
	Period         Period            `json:"period,omitempty"`
	ResourcesUsage *[]TypedResources `json:"resourcesUsage,omitempty"`
}
