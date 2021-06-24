package cloud

// GENERATED SDK for cloud API

// Mongodb cluster definition
type MongodbCreation struct {
	Description  string       `json:"description,omitempty"`
	NetworkId    *string      `json:"networkId,omitempty"`
	NodesList    *[]Node      `json:"nodesList,omitempty"`
	NodesPattern *NodePattern `json:"nodesPattern,omitempty"`
	Plan         string       `json:"plan,omitempty"`
	SubnetId     *string      `json:"subnetId,omitempty"`
	Version      string       `json:"version,omitempty"`
}
