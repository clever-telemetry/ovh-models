package cloud

// GENERATED SDK for cloud API

// Enum values for NodePool Status
type NodePoolStatusEnum string

var (
	NodePoolStatusEnumINSTALLING  NodePoolStatusEnum = "INSTALLING"
	NodePoolStatusEnumUPDATING    NodePoolStatusEnum = "UPDATING"
	NodePoolStatusEnumREDEPLOYING NodePoolStatusEnum = "REDEPLOYING"
	NodePoolStatusEnumRESIZING    NodePoolStatusEnum = "RESIZING"
	NodePoolStatusEnumRESETTING   NodePoolStatusEnum = "RESETTING"
	NodePoolStatusEnumDELETING    NodePoolStatusEnum = "DELETING"
	NodePoolStatusEnumERROR       NodePoolStatusEnum = "ERROR"
	NodePoolStatusEnumREADY       NodePoolStatusEnum = "READY"
)
