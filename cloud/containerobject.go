package cloud

// GENERATED SDK for cloud API

// ContainerObject
type ContainerObject struct {
	ContentType    string             `json:"contentType,omitempty"`
	LastModified   string             `json:"lastModified,omitempty"`
	Name           string             `json:"name,omitempty"`
	RetrievalDelay int64              `json:"retrievalDelay,omitempty"`
	RetrievalState RetrievalStateEnum `json:"retrievalState,omitempty"`
	Size           int64              `json:"size,omitempty"`
}
