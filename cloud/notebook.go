package cloud

// GENERATED SDK for cloud API

// AI Solutions Platform Notebook Object
type Notebook struct {
	CreatedAt string         `json:"createdAt,omitempty"`
	Id        string         `json:"id,omitempty"`
	Spec      NotebookSpec   `json:"spec,omitempty"`
	Status    NotebookStatus `json:"status,omitempty"`
	UpdatedAt string         `json:"updatedAt,omitempty"`
	User      string         `json:"user,omitempty"`
}
