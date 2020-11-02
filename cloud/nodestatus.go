package cloud

// GENERATED SDK for cloud API

// Enum values for Status
type NodeStatus string

var (
	NodeStatusUSER_QUOTA_ERROR            NodeStatus = "USER_QUOTA_ERROR"
	NodeStatusUSER_NODE_SUSPENDED_SERVICE NodeStatus = "USER_NODE_SUSPENDED_SERVICE"
	NodeStatusUSER_NODE_NOT_FOUND_ERROR   NodeStatus = "USER_NODE_NOT_FOUND_ERROR"
	NodeStatusUSER_ERROR                  NodeStatus = "USER_ERROR"
	NodeStatusUPDATING                    NodeStatus = "UPDATING"
	NodeStatusSUSPENDING                  NodeStatus = "SUSPENDING"
	NodeStatusSUSPENDED                   NodeStatus = "SUSPENDED"
	NodeStatusRESETTING                   NodeStatus = "RESETTING"
	NodeStatusREOPENING                   NodeStatus = "REOPENING"
	NodeStatusREADY                       NodeStatus = "READY"
	NodeStatusINSTALLING                  NodeStatus = "INSTALLING"
	NodeStatusERROR                       NodeStatus = "ERROR"
	NodeStatusDELETING                    NodeStatus = "DELETING"
)
