package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type ClusterStatus string

var (
	ClusterStatusUSER_QUOTA_ERROR ClusterStatus = "USER_QUOTA_ERROR"
	ClusterStatusUSER_ERROR       ClusterStatus = "USER_ERROR"
	ClusterStatusUPDATING         ClusterStatus = "UPDATING"
	ClusterStatusSUSPENDING       ClusterStatus = "SUSPENDING"
	ClusterStatusSUSPENDED        ClusterStatus = "SUSPENDED"
	ClusterStatusRESETTING        ClusterStatus = "RESETTING"
	ClusterStatusREOPENING        ClusterStatus = "REOPENING"
	ClusterStatusREADY            ClusterStatus = "READY"
	ClusterStatusINSTALLING       ClusterStatus = "INSTALLING"
	ClusterStatusERROR            ClusterStatus = "ERROR"
	ClusterStatusDELETING         ClusterStatus = "DELETING"
)
