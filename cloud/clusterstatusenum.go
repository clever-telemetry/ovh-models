package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type ClusterStatusEnum string

var (
	ClusterStatusEnumUSER_QUOTA_ERROR            ClusterStatusEnum = "USER_QUOTA_ERROR"
	ClusterStatusEnumUSER_NODE_SUSPENDED_SERVICE ClusterStatusEnum = "USER_NODE_SUSPENDED_SERVICE"
	ClusterStatusEnumUSER_NODE_NOT_FOUND_ERROR   ClusterStatusEnum = "USER_NODE_NOT_FOUND_ERROR"
	ClusterStatusEnumUSER_ERROR                  ClusterStatusEnum = "USER_ERROR"
	ClusterStatusEnumUPDATING                    ClusterStatusEnum = "UPDATING"
	ClusterStatusEnumSUSPENDING                  ClusterStatusEnum = "SUSPENDING"
	ClusterStatusEnumSUSPENDED                   ClusterStatusEnum = "SUSPENDED"
	ClusterStatusEnumRESIZING                    ClusterStatusEnum = "RESIZING"
	ClusterStatusEnumRESETTING                   ClusterStatusEnum = "RESETTING"
	ClusterStatusEnumREOPENING                   ClusterStatusEnum = "REOPENING"
	ClusterStatusEnumREDEPLOYING                 ClusterStatusEnum = "REDEPLOYING"
	ClusterStatusEnumREADY                       ClusterStatusEnum = "READY"
	ClusterStatusEnumMAINTENANCE                 ClusterStatusEnum = "MAINTENANCE"
	ClusterStatusEnumINSTALLING                  ClusterStatusEnum = "INSTALLING"
	ClusterStatusEnumERROR                       ClusterStatusEnum = "ERROR"
	ClusterStatusEnumDELETING                    ClusterStatusEnum = "DELETING"
	ClusterStatusEnumDELETED                     ClusterStatusEnum = "DELETED"
)
