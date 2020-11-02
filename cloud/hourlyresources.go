package cloud

// GENERATED SDK for cloud API

// HourlyResources
type HourlyResources struct {
	Instance          []HourlyInstance          `json:"instance,omitempty"`
	InstanceBandwidth []HourlyInstanceBandwidth `json:"instanceBandwidth,omitempty"`
	InstanceOption    []HourlyInstanceOption    `json:"instanceOption,omitempty"`
	Snapshot          []HourlySnapshot          `json:"snapshot,omitempty"`
	Storage           []HourlyStorage           `json:"storage,omitempty"`
	Volume            []HourlyVolume            `json:"volume,omitempty"`
}
