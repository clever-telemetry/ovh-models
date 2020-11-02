package cloud

// GENERATED SDK for cloud API

// InstanceGroup
type InstanceGroup struct {
	Id           string                `json:"id,omitempty"`
	Instance_ids []string              `json:"instance_ids,omitempty"`
	Name         string                `json:"name,omitempty"`
	Region       string                `json:"region,omitempty"`
	Type         InstanceGroupTypeEnum `json:"type,omitempty"`
}
