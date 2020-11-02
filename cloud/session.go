package cloud

// GENERATED SDK for cloud API

// Session
type Session struct {
	Expires   string  `json:"expires,omitempty"`
	Id        string  `json:"id,omitempty"`
	Profile   Profile `json:"profile,omitempty"`
	Websocket string  `json:"websocket,omitempty"`
}
