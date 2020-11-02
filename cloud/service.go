package cloud

// GENERATED SDK for cloud API

// Details about a Service
type Service struct {
	CanDeleteAtExpiration bool            `json:"canDeleteAtExpiration,omitempty"`
	ContactAdmin          string          `json:"contactAdmin,omitempty"`
	ContactBilling        string          `json:"contactBilling,omitempty"`
	ContactTech           string          `json:"contactTech,omitempty"`
	Creation              string          `json:"creation,omitempty"`
	Domain                string          `json:"domain,omitempty"`
	EngagedUpTo           *string         `json:"engagedUpTo,omitempty"`
	Expiration            string          `json:"expiration,omitempty"`
	PossibleRenewPeriod   *[]int64        `json:"possibleRenewPeriod,omitempty"`
	Renew                 *RenewType      `json:"renew,omitempty"`
	RenewalType           RenewalTypeEnum `json:"renewalType,omitempty"`
	ServiceId             int64           `json:"serviceId,omitempty"`
	Status                StateEnum       `json:"status,omitempty"`
}
