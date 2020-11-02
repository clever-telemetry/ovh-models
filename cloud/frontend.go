package cloud

// GENERATED SDK for cloud API

// A load balancer frontend
type Frontend struct {
	Backends  []BackendSelector `json:"backends"`
	Mode      *ModeEnum         `json:"mode,omitempty"`
	Name      string            `json:"name"`
	Port      int64             `json:"port"`
	Whitelist []string          `json:"whitelist,omitempty"`
}
