package proxy

import (
	"crypto/tls"
	"net/http"
)

// TLSConfig provides basic TLS configuration
func TLSConfig(certFile, keyFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}, nil
}

// HTTPSProxy creates an HTTPS server with the provided handler and TLS configuration
func HTTPSProxy(handler http.Handler, certFile, keyFile string) (*http.Server, error) {
	tlsConfig, err := TLSConfig(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:      ":443",
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	return server, nil
}
