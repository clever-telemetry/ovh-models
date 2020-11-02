package cloud

// GENERATED SDK for cloud API

// Status of a load balancer
type StatusEnum string

var (
	StatusEnumRUNNING  StatusEnum = "RUNNING"
	StatusEnumFROZEN   StatusEnum = "FROZEN"
	StatusEnumERROR    StatusEnum = "ERROR"
	StatusEnumDELETING StatusEnum = "DELETING"
	StatusEnumCREATED  StatusEnum = "CREATED"
	StatusEnumAPPLYING StatusEnum = "APPLYING"
)
