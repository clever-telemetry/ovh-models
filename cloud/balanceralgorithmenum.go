package cloud

// GENERATED SDK for cloud API

// Available load balancer backend balancer algorithm
type BalancerAlgorithmEnum string

var (
	BalancerAlgorithmEnumSTATIC_RR  BalancerAlgorithmEnum = "static-rr"
	BalancerAlgorithmEnumSOURCE     BalancerAlgorithmEnum = "source"
	BalancerAlgorithmEnumROUNDROBIN BalancerAlgorithmEnum = "roundrobin"
	BalancerAlgorithmEnumLEASTCONN  BalancerAlgorithmEnum = "leastconn"
	BalancerAlgorithmEnumFIRST      BalancerAlgorithmEnum = "first"
)
