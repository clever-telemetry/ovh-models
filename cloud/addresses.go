package cloud

// GENERATED SDK for cloud API

// IP list split in version 4 and 6
type Addresses struct {
	Ipv4 []string  `json:"ipv4,omitempty"`
	Ipv6 *[]string `json:"ipv6,omitempty"`
}
