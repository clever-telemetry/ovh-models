package cloud

// GENERATED SDK for cloud API

// Mongodb backup definition
type Backup struct {
	CreatedAt   string      `json:"createdAt,omitempty"`
	Description string      `json:"description,omitempty"`
	Id          string      `json:"id,omitempty"`
	Size        interface{} `json:"size,omitempty"`
	Status      StatusEnum  `json:"status,omitempty"`
}
