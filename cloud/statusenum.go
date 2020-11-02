package cloud

// GENERATED SDK for cloud API

// Status of a load balancer
type StatusEnum string

var (
	StatusEnumCREATED  StatusEnum = "CREATED"
	StatusEnumAPPLYING StatusEnum = "APPLYING"
	StatusEnumRUNNING  StatusEnum = "RUNNING"
	StatusEnumDELETING StatusEnum = "DELETING"
	StatusEnumERROR    StatusEnum = "ERROR"
	StatusEnumFROZEN   StatusEnum = "FROZEN"
)
