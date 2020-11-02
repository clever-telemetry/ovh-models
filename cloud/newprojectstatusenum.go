package cloud

// GENERATED SDK for cloud API

// Possible values for new project status
type NewProjectStatusEnum string

var (
	NewProjectStatusEnumWAITINGAGREEMENTSVALIDATION NewProjectStatusEnum = "waitingAgreementsValidation"
	NewProjectStatusEnumVALIDATIONPENDING           NewProjectStatusEnum = "validationPending"
	NewProjectStatusEnumOK                          NewProjectStatusEnum = "ok"
	NewProjectStatusEnumCREATING                    NewProjectStatusEnum = "creating"
)
