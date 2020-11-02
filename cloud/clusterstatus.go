package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type ClusterStatus string

var (
	ClusterStatusINSTALLING       ClusterStatus = "INSTALLING"
	ClusterStatusUPDATING         ClusterStatus = "UPDATING"
	ClusterStatusRESETTING        ClusterStatus = "RESETTING"
	ClusterStatusSUSPENDING       ClusterStatus = "SUSPENDING"
	ClusterStatusREOPENING        ClusterStatus = "REOPENING"
	ClusterStatusDELETING         ClusterStatus = "DELETING"
	ClusterStatusSUSPENDED        ClusterStatus = "SUSPENDED"
	ClusterStatusERROR            ClusterStatus = "ERROR"
	ClusterStatusUSER_ERROR       ClusterStatus = "USER_ERROR"
	ClusterStatusUSER_QUOTA_ERROR ClusterStatus = "USER_QUOTA_ERROR"
	ClusterStatusREADY            ClusterStatus = "READY"
)
