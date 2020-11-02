package cloud

// GENERATED SDK for cloud API

// List your automated backups
type Backup struct {
	BackupName string       `json:"backupName,omitempty"`
	CreatedAt  string       `json:"createdAt,omitempty"`
	Cron       string       `json:"cron,omitempty"`
	Executions *[]Execution `json:"executions,omitempty"`
	Id         string       `json:"id,omitempty"`
	InstanceId string       `json:"instanceId,omitempty"`
	Name       string       `json:"name,omitempty"`
}
