package cloud

// GENERATED SDK for cloud API

// Status of current version
type VersionStatusEnum string

var (
	VersionStatusEnumROLLBACK    VersionStatusEnum = "rollback"
	VersionStatusEnumPENDING     VersionStatusEnum = "pending"
	VersionStatusEnumFAILED      VersionStatusEnum = "failed"
	VersionStatusEnumDEPLOYING   VersionStatusEnum = "deploying"
	VersionStatusEnumDEPLOYED    VersionStatusEnum = "deployed"
	VersionStatusEnumBUILT       VersionStatusEnum = "built"
	VersionStatusEnumBUILDING    VersionStatusEnum = "building"
	VersionStatusEnumBUILD_ERROR VersionStatusEnum = "build-error"
)
