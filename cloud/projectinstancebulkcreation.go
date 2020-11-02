package cloud

// GENERATED SDK for cloud API

// Missing description
type ProjectInstanceBulkCreation struct {
	Autobackup     *AutoBackup          `json:"autobackup,omitempty"`
	FlavorId       string               `json:"flavorId"`
	GroupId        *string              `json:"groupId,omitempty"`
	ImageId        *string              `json:"imageId,omitempty"`
	MonthlyBilling *bool                `json:"monthlyBilling,omitempty"`
	Name           string               `json:"name"`
	Networks       *[]NetworkBulkParams `json:"networks,omitempty"`
	Number         int64                `json:"number"`
	Region         string               `json:"region"`
	SshKeyId       *string              `json:"sshKeyId,omitempty"`
	UserData       *string              `json:"userData,omitempty"`
	VolumeId       *string              `json:"volumeId,omitempty"`
}
