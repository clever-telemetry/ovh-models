package cloud

// GENERATED SDK for cloud API

// Migration
type Migration struct {
	Date         string           `json:"date,omitempty"`
	MigrationId  string           `json:"migrationId,omitempty"`
	ResourceId   string           `json:"resourceId,omitempty"`
	ResourceType ResourceTypeEnum `json:"resourceType,omitempty"`
}
