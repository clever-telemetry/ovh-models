package cloud

// GENERATED SDK for cloud API

// Status of API
type APIStatusEnum string

var (
	APIStatusEnumPENDING  APIStatusEnum = "pending"
	APIStatusEnumSTARTING APIStatusEnum = "starting"
	APIStatusEnumRUNNING  APIStatusEnum = "running"
	APIStatusEnumSCALING  APIStatusEnum = "scaling"
	APIStatusEnumWAKING   APIStatusEnum = "waking"
	APIStatusEnumSLEEPING APIStatusEnum = "sleeping"
)
