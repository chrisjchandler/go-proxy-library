package proxy

import (
	"net/http"
)

// ReverseProxyRoundRobin creates a reverse proxy with round-robin load balancing
func ReverseProxyRoundRobin(targets []string) (http.Handler, error) {
	return NewRoundRobinProxy(targets)
}
