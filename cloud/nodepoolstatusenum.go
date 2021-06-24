package cloud

// GENERATED SDK for cloud API

// Enum values for NodePool Status
type NodePoolStatusEnum string

var (
	NodePoolStatusEnumUSER_QUOTA_ERROR            NodePoolStatusEnum = "USER_QUOTA_ERROR"
	NodePoolStatusEnumUSER_NODE_SUSPENDED_SERVICE NodePoolStatusEnum = "USER_NODE_SUSPENDED_SERVICE"
	NodePoolStatusEnumUSER_NODE_NOT_FOUND_ERROR   NodePoolStatusEnum = "USER_NODE_NOT_FOUND_ERROR"
	NodePoolStatusEnumUSER_ERROR                  NodePoolStatusEnum = "USER_ERROR"
	NodePoolStatusEnumUPDATING                    NodePoolStatusEnum = "UPDATING"
	NodePoolStatusEnumSUSPENDING                  NodePoolStatusEnum = "SUSPENDING"
	NodePoolStatusEnumSUSPENDED                   NodePoolStatusEnum = "SUSPENDED"
	NodePoolStatusEnumRESIZING                    NodePoolStatusEnum = "RESIZING"
	NodePoolStatusEnumRESETTING                   NodePoolStatusEnum = "RESETTING"
	NodePoolStatusEnumREOPENING                   NodePoolStatusEnum = "REOPENING"
	NodePoolStatusEnumREDEPLOYING                 NodePoolStatusEnum = "REDEPLOYING"
	NodePoolStatusEnumREADY                       NodePoolStatusEnum = "READY"
	NodePoolStatusEnumMAINTENANCE                 NodePoolStatusEnum = "MAINTENANCE"
	NodePoolStatusEnumINSTALLING                  NodePoolStatusEnum = "INSTALLING"
	NodePoolStatusEnumERROR                       NodePoolStatusEnum = "ERROR"
	NodePoolStatusEnumDELETING                    NodePoolStatusEnum = "DELETING"
	NodePoolStatusEnumDELETED                     NodePoolStatusEnum = "DELETED"
)
