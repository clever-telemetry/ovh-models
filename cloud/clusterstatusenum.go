package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type ClusterStatusEnum string

var (
	ClusterStatusEnumINSTALLING       ClusterStatusEnum = "INSTALLING"
	ClusterStatusEnumUPDATING         ClusterStatusEnum = "UPDATING"
	ClusterStatusEnumREDEPLOYING      ClusterStatusEnum = "REDEPLOYING"
	ClusterStatusEnumRESETTING        ClusterStatusEnum = "RESETTING"
	ClusterStatusEnumSUSPENDING       ClusterStatusEnum = "SUSPENDING"
	ClusterStatusEnumREOPENING        ClusterStatusEnum = "REOPENING"
	ClusterStatusEnumDELETING         ClusterStatusEnum = "DELETING"
	ClusterStatusEnumSUSPENDED        ClusterStatusEnum = "SUSPENDED"
	ClusterStatusEnumMAINTENANCE      ClusterStatusEnum = "MAINTENANCE"
	ClusterStatusEnumERROR            ClusterStatusEnum = "ERROR"
	ClusterStatusEnumUSER_ERROR       ClusterStatusEnum = "USER_ERROR"
	ClusterStatusEnumUSER_QUOTA_ERROR ClusterStatusEnum = "USER_QUOTA_ERROR"
	ClusterStatusEnumREADY            ClusterStatusEnum = "READY"
)
