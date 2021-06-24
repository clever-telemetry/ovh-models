package cloud

// GENERATED SDK for cloud API

// Job information
type Job struct {
	ContainerName    string            `json:"containerName"`
	CreationDate     *string           `json:"creationDate,omitempty"`
	EndDate          *string           `json:"endDate,omitempty"`
	Engine           string            `json:"engine"`
	EngineParameters []EngineParameter `json:"engineParameters"`
	EngineVersion    string            `json:"engineVersion"`
	Id               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Region           string            `json:"region"`
	StartDate        *string           `json:"startDate,omitempty"`
	Status           StatusEnum        `json:"status,omitempty"`
	Ttl              *string           `json:"ttl,omitempty"`
}
