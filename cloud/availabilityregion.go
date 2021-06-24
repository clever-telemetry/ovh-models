package cloud

// GENERATED SDK for cloud API

// Details about a region
type AvailabilityRegion struct {
	ContinentCode RegionContinentEnum `json:"continentCode,omitempty"`
	Datacenter    string              `json:"datacenter,omitempty"`
	Enabled       bool                `json:"enabled,omitempty"`
	Name          string              `json:"name,omitempty"`
}
