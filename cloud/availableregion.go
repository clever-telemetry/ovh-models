package cloud

// GENERATED SDK for cloud API

// Details about an available region that can be activated on your project
type AvailableRegion struct {
	ContinentCode      RegionContinentEnum `json:"continentCode,omitempty"`
	DatacenterLocation string              `json:"datacenterLocation,omitempty"`
	Name               string              `json:"name,omitempty"`
}
