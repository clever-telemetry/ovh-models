package cloud

// GENERATED SDK for cloud API

// Availability of databases engines on cloud projects
type Availability struct {
	Backup            BackupTypeEnum  `json:"backup,omitempty"`
	Default           bool            `json:"default,omitempty"`
	EndOfLife         *string         `json:"endOfLife,omitempty"`
	Engine            string          `json:"engine,omitempty"`
	Flavor            string          `json:"flavor,omitempty"`
	MaxNodeNumber     int64           `json:"maxNodeNumber,omitempty"`
	MinNodeNumber     int64           `json:"minNodeNumber,omitempty"`
	Network           NetworkTypeEnum `json:"network,omitempty"`
	Plan              string          `json:"plan,omitempty"`
	Region            string          `json:"region,omitempty"`
	StartDate         string          `json:"startDate,omitempty"`
	UpstreamEndOfLife *string         `json:"upstreamEndOfLife,omitempty"`
	Version           string          `json:"version,omitempty"`
}
