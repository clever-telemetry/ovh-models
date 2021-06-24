package cloud

// GENERATED SDK for cloud API

// Add storage policy for container
type AddContainerPolicy struct {
	ObjectKey string         `json:"objectKey,omitempty"`
	RoleName  PolicyRoleEnum `json:"roleName"`
}
