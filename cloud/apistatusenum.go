package cloud

// GENERATED SDK for cloud API

// Status of API
type APIStatusEnum string

var (
	APIStatusEnumWAKING   APIStatusEnum = "waking"
	APIStatusEnumSTARTING APIStatusEnum = "starting"
	APIStatusEnumSLEEPING APIStatusEnum = "sleeping"
	APIStatusEnumSCALING  APIStatusEnum = "scaling"
	APIStatusEnumRUNNING  APIStatusEnum = "running"
	APIStatusEnumPENDING  APIStatusEnum = "pending"
)
