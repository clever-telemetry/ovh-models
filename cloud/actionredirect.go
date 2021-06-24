package cloud

// GENERATED SDK for cloud API

// HTTP load balancer redirect action
type ActionRedirect struct {
	Location   string                 `json:"location"`
	Name       string                 `json:"name"`
	StatusCode RedirectStatusCodeEnum `json:"statusCode"`
}
