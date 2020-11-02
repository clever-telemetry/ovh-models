package cloud

// GENERATED SDK for cloud API

// Autobackup params at instance creation
type AutoBackup struct {
	Cron     string `json:"cron,omitempty"`
	Rotation int64  `json:"rotation,omitempty"`
}
