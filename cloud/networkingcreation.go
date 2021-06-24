package cloud

// GENERATED SDK for cloud API

// Networking creation object
type NetworkingCreation struct {
	Egress  *EgressCreation  `json:"egress,omitempty"`
	Ingress *IngressCreation `json:"ingress,omitempty"`
}
