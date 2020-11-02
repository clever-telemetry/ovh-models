package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectUserCreation struct {
	Description *string     `json:"description,omitempty"`
	Role        *RoleEnum   `json:"role,omitempty"`
	Roles       *[]RoleEnum `json:"roles,omitempty"`
}
