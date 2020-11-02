package cloud

// GENERATED SDK for cloud API

// Status of current version
type VersionStatusEnum string

var (
	VersionStatusEnumPENDING     VersionStatusEnum = "pending"
	VersionStatusEnumBUILDING    VersionStatusEnum = "building"
	VersionStatusEnumBUILT       VersionStatusEnum = "built"
	VersionStatusEnumBUILD_ERROR VersionStatusEnum = "build-error"
	VersionStatusEnumDEPLOYING   VersionStatusEnum = "deploying"
	VersionStatusEnumDEPLOYED    VersionStatusEnum = "deployed"
	VersionStatusEnumROLLBACK    VersionStatusEnum = "rollback"
	VersionStatusEnumFAILED      VersionStatusEnum = "failed"
)
