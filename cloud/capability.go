package cloud

// GENERATED SDK for cloud API

// Capabilities of data processing service
type Capability struct {
	AvailableVersions []EngineVersion               `json:"availableVersions,omitempty"`
	Name              string                        `json:"name,omitempty"`
	Parameters        []CapabilitiesEngineParameter `json:"parameters,omitempty"`
	Templates         []CapabilitiesTemplate        `json:"templates,omitempty"`
}
