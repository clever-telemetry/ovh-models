package cloud

// GENERATED SDK for cloud API

// Mongodb node definition
type Node struct {
	CreatedAt string     `json:"createdAt,omitempty"`
	Flavor    string     `json:"flavor,omitempty"`
	Id        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Port      int64      `json:"port,omitempty"`
	Region    string     `json:"region,omitempty"`
	Status    StatusEnum `json:"status,omitempty"`
}
