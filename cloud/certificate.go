package cloud

// GENERATED SDK for cloud API

// A Certificate to use in your NFVs
type Certificate struct {
	ExpireAt               string                  `json:"expireAt,omitempty"`
	Fingerprint            string                  `json:"fingerprint,omitempty"`
	Id                     string                  `json:"id,omitempty"`
	Issuer                 string                  `json:"issuer,omitempty"`
	Kind                   CertificateKindEnum     `json:"kind,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	SerialNumber           string                  `json:"serialNumber,omitempty"`
	ServerAlternativeNames []ServerAlternativeName `json:"serverAlternativeNames,omitempty"`
	Status                 CertificateStatusEnum   `json:"status,omitempty"`
	Subject                string                  `json:"subject,omitempty"`
	ValidAt                string                  `json:"validAt,omitempty"`
	Version                int64                   `json:"version,omitempty"`
}
