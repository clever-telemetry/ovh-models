package cloud

// GENERATED SDK for cloud API

// HTTP load balancer actions
type Actions struct {
	Dispatch *[]ActionDispatch `json:"dispatch,omitempty"`
	Redirect *[]ActionRedirect `json:"redirect,omitempty"`
	Reject   *[]ActionReject   `json:"reject,omitempty"`
	Rewrite  *[]ActionRewrite  `json:"rewrite,omitempty"`
}
