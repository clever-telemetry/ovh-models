package cloud

// GENERATED SDK for cloud API

// A Image of a built serving model
type PresetImage struct {
	Description      string  `json:"description,omitempty"`
	Id               string  `json:"id,omitempty"`
	Link             *string `json:"link,omitempty"`
	Name             string  `json:"name,omitempty"`
	RamRequirementMB *int64  `json:"ramRequirementMB,omitempty"`
}
