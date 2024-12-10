# Go Proxy Library



A Go-based library for building forward and reverse proxies with support for round-robin load balancing and SSL/TLS termination. This library is designed to be lightweight, flexible, and easy to integrate into your applications.

## Features

- **Forward Proxy:** Route outbound requests to multiple upstream servers.
- **Reverse Proxy:** Handle inbound client requests and forward them to backend servers.
- **Round-Robin Load Balancing:** Distribute requests across multiple servers in a round-robin fashion.
- **SSL/TLS Termination:** Secure your proxy with HTTPS.
- **Configurable:** Easily configure proxy modes, backend servers, and SSL settings.

Configuration: Edit the proxy settings in the config struct

API Reference
Forward Proxy
func ForwardProxyRoundRobin(targets []string) (http.Handler, error)
Creates a forward proxy with round-robin load balancing.
Reverse Proxy
func ReverseProxyRoundRobin(targets []string) (http.Handler, error)
Creates a reverse proxy with round-robin load balancing.
SSL/TLS Support
func TLSConfig(certFile, keyFile string) (*tls.Config, error)
Generates a TLS configuration using provided certificate and key files.
func HTTPSProxy(handler http.Handler, certFile, keyFile string) (*http.Server, error)
Creates an HTTPS server with TLS support.

A Go-based library for building forward and reverse proxies with support for round-robin load balancing and SSL/TLS termination. This library is designed to be lightweight, flexible, and easy to integrate into your applications.

## Features

- **Forward Proxy:** Route outbound requests to multiple upstream servers.
- **Reverse Proxy:** Handle inbound client requests and forward them to backend servers.
- **Round-Robin Load Balancing:** Distribute requests across multiple servers in a round-robin fashion.
- **SSL/TLS Termination:** Secure your proxy with HTTPS.
- **Configurable:** Easily configure proxy modes, backend servers, and SSL settings.

---


Future Enhancements
Planned features include:

Health checks for backend servers.
Sticky sessions for client-specific routing.
Dynamic addition/removal of backends.
Weighted round-robin for traffic distribution.
