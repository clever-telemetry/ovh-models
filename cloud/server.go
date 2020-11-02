package cloud

// GENERATED SDK for cloud API

// A load balancer backend server
type Server struct {
	Ip      string `json:"ip"`
	Name    string `json:"name"`
	NoCheck *bool  `json:"noCheck,omitempty"`
	Port    int64  `json:"port"`
	Weight  *int64 `json:"weight,omitempty"`
}
