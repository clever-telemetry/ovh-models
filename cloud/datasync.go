package cloud

// GENERATED SDK for cloud API

// AI Solutions Data Sync
type DataSync struct {
	CreatedAt string         `json:"createdAt,omitempty"`
	Id        string         `json:"id,omitempty"`
	Spec      DataSyncSpec   `json:"spec,omitempty"`
	Status    DataSyncStatus `json:"status,omitempty"`
	UpdatedAt string         `json:"updatedAt,omitempty"`
}
