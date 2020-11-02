package cloud

// GENERATED SDK for cloud API

// State of the job
type JobStateEnum string

var (
	JobStateEnumQUEUED      JobStateEnum = "QUEUED"
	JobStateEnumPENDING     JobStateEnum = "PENDING"
	JobStateEnumSYNCING     JobStateEnum = "SYNCING"
	JobStateEnumRUNNING     JobStateEnum = "RUNNING"
	JobStateEnumFAILED      JobStateEnum = "FAILED"
	JobStateEnumERROR       JobStateEnum = "ERROR"
	JobStateEnumDONE        JobStateEnum = "DONE"
	JobStateEnumINTERRUPTED JobStateEnum = "INTERRUPTED"
)
