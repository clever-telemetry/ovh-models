package cloud

// GENERATED SDK for cloud API

// FailoverIp
type FailoverIp struct {
	Block         *string       `json:"block,omitempty"`
	ContinentCode *string       `json:"continentCode,omitempty"`
	Geoloc        *string       `json:"geoloc,omitempty"`
	Id            string        `json:"id,omitempty"`
	Ip            *string       `json:"ip,omitempty"`
	Progress      int64         `json:"progress,omitempty"`
	RoutedTo      string        `json:"routedTo,omitempty"`
	Status        IpStatusEnum  `json:"status,omitempty"`
	SubType       IpSubTypeEnum `json:"subType,omitempty"`
}
