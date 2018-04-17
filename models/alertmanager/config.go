package alertmanager

import (
	"regexp"
)

// Secret represents a Secret string
type Secret string

// Regexp represents a Regexp
type Regexp struct {
	*regexp.Regexp
}

// BasicAuth defines Basic Auth information
type BasicAuth struct {
	Username string
	Password Secret
}

// TLSConfig defines TLS configuration items
type TLSConfig struct {
	// CA certificate to validate the server certificate with.
	CAFile string
	// Certificate and key files for client cert authentication to the server.
	CertFile string
	KeyFile  string
	// ServerName extension to indicate the name of the server.
	ServerName string
	// Disable validation of the server certificate.
	InsecureSkipVerity bool
}

// HTTPConfig defines essential HTTP communication settings
type HTTPConfig struct {
	// Note that `basic_auth`, `bearer_token` and `bearer_token_file` options are mutually exclusive.
	// Sets the `Authorization` header with the configured username and password.
	BasicAuth *BasicAuth
	// Sets the `Authorization` header with the configured bearer token.
	BearerToken Secret
	// Sets the `Authorization` header with the bearer token read from the configured file.
	BearerTokenFile string
	// Configures the TLS settings.
	TLSConfig *TLSConfig
	// Optional proxy URL.
	ProxyURL string
}

// Config encapsulates alertmanager configuration file
type Config struct {
	Global       *Global
	Route        *Route
	InhibitRules []*InhibitRule
	Receivers    []*Receiver
	Templates    []string
}
