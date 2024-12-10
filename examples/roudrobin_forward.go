package main

import (
	"log"
	"net/http"

	"github.com/chrisjchandler/go-proxy-library/proxy"
)

func main() {
	// List of backend servers for forward proxy
	backends := []string{
		"https://server1.foo.com",
		"https://server2.bar.com",
		"https://server3.foobar.com",
	}

	// Create a round-robin forward proxy
	handler, err := proxy.ForwardProxyRoundRobin(backends)
	if err != nil {
		log.Fatalf("Failed to create forward proxy: %v", err)
	}

	// Start the proxy server
	port := ":8080"
	log.Printf("Starting forward proxy with round-robin on port %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
