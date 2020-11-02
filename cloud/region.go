package cloud

// GENERATED SDK for cloud API

// Details about your region
type Region struct {
	ContinentCode      RegionContinentEnum `json:"continentCode,omitempty"`
	DatacenterLocation string              `json:"datacenterLocation,omitempty"`
	IpCountries        []IpCountryEnum     `json:"ipCountries,omitempty"`
	Name               string              `json:"name,omitempty"`
	Services           []Component         `json:"services,omitempty"`
	Status             RegionStatusEnum    `json:"status,omitempty"`
}
