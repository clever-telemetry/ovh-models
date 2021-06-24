package cloud

// GENERATED SDK for cloud API

// An application load balancer configuration
type ApplicationConfigurationCreation struct {
	Actions      *Actions     `json:"actions,omitempty"`
	Certificates []string     `json:"certificates,omitempty"`
	Conditions   *[]Condition `json:"conditions,omitempty"`
	EntryPoints  []EntryPoint `json:"entryPoints,omitempty"`
	Networking   *Networking  `json:"networking,omitempty"`
	Targets      *[]Target    `json:"targets,omitempty"`
	Version      int64        `json:"version,omitempty"`
}
