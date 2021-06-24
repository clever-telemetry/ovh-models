package cloud

// GENERATED SDK for cloud API

// State of the data sync
type DataSyncStateEnum string

var (
	DataSyncStateEnumRUNNING DataSyncStateEnum = "RUNNING"
	DataSyncStateEnumQUEUED  DataSyncStateEnum = "QUEUED"
	DataSyncStateEnumFAILED  DataSyncStateEnum = "FAILED"
	DataSyncStateEnumERROR   DataSyncStateEnum = "ERROR"
	DataSyncStateEnumDONE    DataSyncStateEnum = "DONE"
)
