package cloud

// GENERATED SDK for cloud API

// User
type User struct {
	CreationDate string         `json:"creationDate,omitempty"`
	Description  string         `json:"description,omitempty"`
	Id           int64          `json:"id,omitempty"`
	OpenstackId  *string        `json:"openstackId,omitempty"`
	Roles        []Role         `json:"roles,omitempty"`
	Status       UserStatusEnum `json:"status,omitempty"`
	Username     string         `json:"username,omitempty"`
}
