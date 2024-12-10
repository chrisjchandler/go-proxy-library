package proxy

// Config represents the configuration for the proxy server
type Config struct {
	Mode       string   // "forward" or "reverse"
	Targets    []string // List of backend URLs
	CertFile   string   // Path to the SSL certificate
	KeyFile    string   // Path to the SSL private key
	Port       string   // Port to run the proxy on
}
