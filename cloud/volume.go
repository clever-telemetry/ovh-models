package cloud

// GENERATED SDK for cloud API

// Volume
type Volume struct {
	AttachedTo   []string       `json:"attachedTo,omitempty"`
	Bootable     bool           `json:"bootable,omitempty"`
	CreationDate string         `json:"creationDate,omitempty"`
	Description  string         `json:"description,omitempty"`
	Id           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	PlanCode     *string        `json:"planCode,omitempty"`
	Region       string         `json:"region,omitempty"`
	Size         int64          `json:"size,omitempty"`
	Status       string         `json:"status,omitempty"`
	Type         VolumeTypeEnum `json:"type,omitempty"`
}
