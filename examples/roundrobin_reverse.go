package main

import (
	"log"
	"net/http"

	"github.com/chrisjchandler/go-proxy-library/proxy"
)

func main() {
	// List of backend servers for reverse proxy
	backends := []string{
		"http://foo.com",
		"http://bar.com",
		"http://foobar.com",
	}

	// Create a round-robin reverse proxy
	handler, err := proxy.ReverseProxyRoundRobin(backends)
	if err != nil {
		log.Fatalf("Failed to create reverse proxy: %v", err)
	}

	// Enable SSL (optional)
	certFile := "cert.pem"
	keyFile := "key.pem"
	server, err := proxy.HTTPSProxy(handler, certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to create HTTPS reverse proxy: %v", err)
	}

	// Start the server
	log.Printf("Starting reverse proxy with round-robin on port 443")
	log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
}
