package cloud

// GENERATED SDK for cloud API

// Enum values for NodePool Status
type NodePoolStatusEnum string

var (
	NodePoolStatusEnumUPDATING    NodePoolStatusEnum = "UPDATING"
	NodePoolStatusEnumRESIZING    NodePoolStatusEnum = "RESIZING"
	NodePoolStatusEnumRESETTING   NodePoolStatusEnum = "RESETTING"
	NodePoolStatusEnumREDEPLOYING NodePoolStatusEnum = "REDEPLOYING"
	NodePoolStatusEnumREADY       NodePoolStatusEnum = "READY"
	NodePoolStatusEnumINSTALLING  NodePoolStatusEnum = "INSTALLING"
	NodePoolStatusEnumERROR       NodePoolStatusEnum = "ERROR"
	NodePoolStatusEnumDELETING    NodePoolStatusEnum = "DELETING"
)
