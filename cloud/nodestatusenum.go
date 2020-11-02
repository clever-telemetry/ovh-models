package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type NodeStatusEnum string

var (
	NodeStatusEnumINSTALLING                  NodeStatusEnum = "INSTALLING"
	NodeStatusEnumREDEPLOYING                 NodeStatusEnum = "REDEPLOYING"
	NodeStatusEnumUPDATING                    NodeStatusEnum = "UPDATING"
	NodeStatusEnumRESETTING                   NodeStatusEnum = "RESETTING"
	NodeStatusEnumSUSPENDING                  NodeStatusEnum = "SUSPENDING"
	NodeStatusEnumREOPENING                   NodeStatusEnum = "REOPENING"
	NodeStatusEnumDELETING                    NodeStatusEnum = "DELETING"
	NodeStatusEnumSUSPENDED                   NodeStatusEnum = "SUSPENDED"
	NodeStatusEnumERROR                       NodeStatusEnum = "ERROR"
	NodeStatusEnumUSER_ERROR                  NodeStatusEnum = "USER_ERROR"
	NodeStatusEnumUSER_QUOTA_ERROR            NodeStatusEnum = "USER_QUOTA_ERROR"
	NodeStatusEnumUSER_NODE_NOT_FOUND_ERROR   NodeStatusEnum = "USER_NODE_NOT_FOUND_ERROR"
	NodeStatusEnumUSER_NODE_SUSPENDED_SERVICE NodeStatusEnum = "USER_NODE_SUSPENDED_SERVICE"
	NodeStatusEnumREADY                       NodeStatusEnum = "READY"
)
