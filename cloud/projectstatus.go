package cloud

// GENERATED SDK for cloud API

// Possible values for project status
type ProjectStatus string

var (
	ProjectStatusSUSPENDED ProjectStatus = "suspended"
	ProjectStatusOK        ProjectStatus = "ok"
	ProjectStatusDELETING  ProjectStatus = "deleting"
	ProjectStatusDELETED   ProjectStatus = "deleted"
	ProjectStatusCREATING  ProjectStatus = "creating"
)
