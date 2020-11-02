package cloud

// GENERATED SDK for cloud API

// CloudIp
type CloudIp struct {
	Id     string       `json:"id,omitempty"`
	Ip     *string      `json:"ip,omitempty"`
	Status IpStatusEnum `json:"status,omitempty"`
	Type   string       `json:"type,omitempty"`
}
