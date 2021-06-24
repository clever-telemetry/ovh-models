package cloud

// GENERATED SDK for cloud API

// Mongodb cluster definition
type Mongodb struct {
	CreatedAt         string            `json:"createdAt,omitempty"`
	Description       string            `json:"description,omitempty"`
	Domain            string            `json:"domain,omitempty"`
	Id                string            `json:"id,omitempty"`
	MaintenanceWindow MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	NetworkId         *string           `json:"networkId,omitempty"`
	NetworkType       NetworkTypeEnum   `json:"networkType,omitempty"`
	NodeNumber        int64             `json:"nodeNumber,omitempty"`
	Plan              string            `json:"plan,omitempty"`
	Status            StatusEnum        `json:"status,omitempty"`
	SubnetId          *string           `json:"subnetId,omitempty"`
	Version           string            `json:"version,omitempty"`
}
