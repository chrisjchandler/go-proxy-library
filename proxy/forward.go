package proxy

import (
	"net/http"
)

// ForwardProxyRoundRobin creates a forward proxy with round-robin load balancing
func ForwardProxyRoundRobin(targets []string) (http.Handler, error) {
	return NewRoundRobinProxy(targets)
}
