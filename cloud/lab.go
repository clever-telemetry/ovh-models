package cloud

// GENERATED SDK for cloud API

// A public cloud lab permits to activate a feature in beta
type Lab struct {
	Id     string        `json:"id,omitempty"`
	Name   string        `json:"name,omitempty"`
	Status LabStatusEnum `json:"status,omitempty"`
}
