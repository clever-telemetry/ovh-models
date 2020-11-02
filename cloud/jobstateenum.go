package cloud

// GENERATED SDK for cloud API

// State of the job
type JobStateEnum string

var (
	JobStateEnumRUNNING      JobStateEnum = "RUNNING"
	JobStateEnumQUEUED       JobStateEnum = "QUEUED"
	JobStateEnumPENDING      JobStateEnum = "PENDING"
	JobStateEnumINTERRUPTING JobStateEnum = "INTERRUPTING"
	JobStateEnumINTERRUPTED  JobStateEnum = "INTERRUPTED"
	JobStateEnumINITIALIZING JobStateEnum = "INITIALIZING"
	JobStateEnumFINALIZING   JobStateEnum = "FINALIZING"
	JobStateEnumFAILED       JobStateEnum = "FAILED"
	JobStateEnumERROR        JobStateEnum = "ERROR"
	JobStateEnumDONE         JobStateEnum = "DONE"
)
