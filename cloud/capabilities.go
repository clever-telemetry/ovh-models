package cloud

// GENERATED SDK for cloud API

// Capabilities available for the databases engines on cloud projects
type Capabilities struct {
	Engines []Engine `json:"engines,omitempty"`
	Flavors []Flavor `json:"flavors,omitempty"`
	Options []Option `json:"options,omitempty"`
	Plans   []Plan   `json:"plans,omitempty"`
}
