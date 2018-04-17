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
	Username string `json:"username"`
	Password Secret `json:"password"`
}

// TLSConfig defines TLS configuration items
type TLSConfig struct {
	// CA certificate to validate the server certificate with.
	CAFile string `json:"ca_file"`
	// Certificate and key files for client cert authentication to the server.
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
	// ServerName extension to indicate the name of the server.
	ServerName string `json:"server_name"`
	// Disable validation of the server certificate.
	InsecureSkipVerity bool `json:"insecure_skip_verity"`
}

// HTTPConfig defines essential HTTP communication settings
type HTTPConfig struct {
	// Note that `basic_auth`, `bearer_token` and `bearer_token_file` options are mutually exclusive.
	// Sets the `Authorization` header with the configured username and password.
	BasicAuth *BasicAuth `json:"basic_auth"`
	// Sets the `Authorization` header with the configured bearer token.
	BearerToken Secret `json:"bearer_token"`
	// Sets the `Authorization` header with the bearer token read from the configured file.
	BearerTokenFile string `json:"bearer_token_file"`
	// Configures the TLS settings.
	TLSConfig *TLSConfig `json:"tls_config"`
	// Optional proxy URL.
	ProxyURL string `json:"proxy_url"`
}

// Config encapsulates alertmanager configuration file
type Config struct {
	Global       *Global        `json:"global,omitempty"`
	Templates    []string       `json:"templates"`
	Route        *Route         `json:"route,omitempty"`
	Receivers    []*Receiver    `json:"receivers,omitempty"`
	InhibitRules []*InhibitRule `json:"inhibit_rules,omitempty"`

	// The raw data
	raw string
}
