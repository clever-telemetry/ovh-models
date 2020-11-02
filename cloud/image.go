package cloud

// GENERATED SDK for cloud API

// Image
type Image struct {
	CreationDate string    `json:"creationDate,omitempty"`
	FlavorType   *string   `json:"flavorType,omitempty"`
	Id           string    `json:"id,omitempty"`
	MinDisk      int64     `json:"minDisk,omitempty"`
	MinRam       int64     `json:"minRam,omitempty"`
	Name         string    `json:"name,omitempty"`
	PlanCode     *string   `json:"planCode,omitempty"`
	Region       string    `json:"region,omitempty"`
	Size         float64   `json:"size,omitempty"`
	Status       string    `json:"status,omitempty"`
	Tags         *[]string `json:"tags,omitempty"`
	Type         string    `json:"type,omitempty"`
	User         string    `json:"user,omitempty"`
	Visibility   string    `json:"visibility,omitempty"`
}
