package cloud

// GENERATED SDK for cloud API

// Token to access a stream
type Token struct {
	Action TokenActionEnum `json:"action,omitempty"`
	Id     string          `json:"id,omitempty"`
	Token  string          `json:"token,omitempty"`
}
