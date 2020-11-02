package cloud

// GENERATED SDK for cloud API

// ApplicationAccess
type ApplicationAccess struct {
	Accesses []Access                   `json:"accesses,omitempty"`
	Status   ApplicationAccessStateEnum `json:"status,omitempty"`
}
