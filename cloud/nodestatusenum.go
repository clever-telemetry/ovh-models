package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type NodeStatusEnum string

var (
	NodeStatusEnumUSER_QUOTA_ERROR            NodeStatusEnum = "USER_QUOTA_ERROR"
	NodeStatusEnumUSER_NODE_SUSPENDED_SERVICE NodeStatusEnum = "USER_NODE_SUSPENDED_SERVICE"
	NodeStatusEnumUSER_NODE_NOT_FOUND_ERROR   NodeStatusEnum = "USER_NODE_NOT_FOUND_ERROR"
	NodeStatusEnumUSER_ERROR                  NodeStatusEnum = "USER_ERROR"
	NodeStatusEnumUPDATING                    NodeStatusEnum = "UPDATING"
	NodeStatusEnumSUSPENDING                  NodeStatusEnum = "SUSPENDING"
	NodeStatusEnumSUSPENDED                   NodeStatusEnum = "SUSPENDED"
	NodeStatusEnumRESIZING                    NodeStatusEnum = "RESIZING"
	NodeStatusEnumRESETTING                   NodeStatusEnum = "RESETTING"
	NodeStatusEnumREOPENING                   NodeStatusEnum = "REOPENING"
	NodeStatusEnumREDEPLOYING                 NodeStatusEnum = "REDEPLOYING"
	NodeStatusEnumREADY                       NodeStatusEnum = "READY"
	NodeStatusEnumMAINTENANCE                 NodeStatusEnum = "MAINTENANCE"
	NodeStatusEnumINSTALLING                  NodeStatusEnum = "INSTALLING"
	NodeStatusEnumERROR                       NodeStatusEnum = "ERROR"
	NodeStatusEnumDELETING                    NodeStatusEnum = "DELETING"
	NodeStatusEnumDELETED                     NodeStatusEnum = "DELETED"
)
