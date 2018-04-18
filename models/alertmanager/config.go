package alertmanager

import "github.com/joshuakwan/almond/models/common"

// HTTPConfig defines essential HTTP communication settings
type HTTPConfig struct {
	// Note that `basic_auth`, `bearer_token` and `bearer_token_file` options are mutually exclusive.
	// Sets the `Authorization` header with the configured username and password.
	BasicAuth *common.BasicAuth `json:"basic_auth" yaml:"basic_auth"`
	// Sets the `Authorization` header with the configured bearer token.
	BearerToken common.Secret `json:"bearer_token" yaml:"bearer_token" `
	// Sets the `Authorization` header with the bearer token read from the configured file.
	BearerTokenFile string `json:"bearer_token_file" yaml:"bearer_token_file"`
	// Configures the TLS settings.
	TLSConfig *common.TLSConfig `json:"tls_config" yaml:"tls_config"`
	// Optional proxy URL.
	ProxyURL string `json:"proxy_url" yaml:"proxy_url"`
}

// Config encapsulates alertmanager configuration file
type Config struct {
	Global       *Global        `json:"global,omitempty" yaml:"global,omitempty"`
	Templates    []string       `json:"templates" yaml:"templates"`
	Route        *Route         `json:"route,omitempty" yaml:"route,omitempty"`
	Receivers    []*Receiver    `json:"receivers,omitempty" yaml:"receivers,omitempty"`
	InhibitRules []*InhibitRule `json:"inhibit_rules,omitempty" yaml:"inhibit_rules,omitempty"`

	// The raw data
	raw string
}
