package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

// RoundRobinProxy manages a list of backend servers for round-robin load balancing
type RoundRobinProxy struct {
	backends []*url.URL
	counter  uint32
}

// NewRoundRobinProxy initializes a round-robin proxy with a list of backends
func NewRoundRobinProxy(targets []string) (*RoundRobinProxy, error) {
	var urls []*url.URL
	for _, target := range targets {
		parsed, err := url.Parse(target)
		if err != nil {
			return nil, err
		}
		urls = append(urls, parsed)
	}
	return &RoundRobinProxy{backends: urls}, nil
}

// NextBackend selects the next backend URL in a round-robin manner
func (r *RoundRobinProxy) NextBackend() *url.URL {
	idx := atomic.AddUint32(&r.counter, 1) % uint32(len(r.backends))
	return r.backends[idx]
}

// ServeHTTP handles HTTP requests and forwards them to the next backend
func (r *RoundRobinProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	backend := r.NextBackend()

	// Modify request to match the selected backend
	proxy := httputil.NewSingleHostReverseProxy(backend)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = backend.Scheme
		req.URL.Host = backend.Host
		req.URL.Path = backend.Path
	}
	proxy.ServeHTTP(w, req)
}
