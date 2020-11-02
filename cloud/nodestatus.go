package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type NodeStatus string

var (
	NodeStatusINSTALLING                  NodeStatus = "INSTALLING"
	NodeStatusUPDATING                    NodeStatus = "UPDATING"
	NodeStatusRESETTING                   NodeStatus = "RESETTING"
	NodeStatusSUSPENDING                  NodeStatus = "SUSPENDING"
	NodeStatusREOPENING                   NodeStatus = "REOPENING"
	NodeStatusDELETING                    NodeStatus = "DELETING"
	NodeStatusSUSPENDED                   NodeStatus = "SUSPENDED"
	NodeStatusERROR                       NodeStatus = "ERROR"
	NodeStatusUSER_ERROR                  NodeStatus = "USER_ERROR"
	NodeStatusUSER_QUOTA_ERROR            NodeStatus = "USER_QUOTA_ERROR"
	NodeStatusUSER_NODE_NOT_FOUND_ERROR   NodeStatus = "USER_NODE_NOT_FOUND_ERROR"
	NodeStatusUSER_NODE_SUSPENDED_SERVICE NodeStatus = "USER_NODE_SUSPENDED_SERVICE"
	NodeStatusREADY                       NodeStatus = "READY"
)
