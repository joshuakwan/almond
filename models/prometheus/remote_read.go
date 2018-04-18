package prometheus

import (
	"github.com/joshuakwan/almond/models/common"
	"github.com/prometheus/common/model"
)

// RemoteRead represents remote_read section
type RemoteRead struct {
	// The URL of the endpoint to query from.
	URL string `json:"url" yaml:"url"`

	// An optional list of equality matchers which have to be
	// present in a selector to query the remote read endpoint.
	RequiredMatchers []string `json:"required_matchers,omitempty" yaml:"required_matchers,omitempty"`

	// Timeout for requests to the remote read endpoint.
	RemoteTimeout model.Duration `json:"remote_timeout,omitempty" yaml:"remote_timeout,omitempty"` // default = 1m

	// Whether reads should be made for queries for time ranges that
	// the local storage should have complete data for.
	ReadRecent bool `json:"read_recent,omitempty" yaml:"read_recent,omitempty"` // default = false

	// Sets the `Authorization` header on every scrape request with the
	// configured username and password.
	BasicAuth *common.BasicAuth `json:"basic_auth,omitempty" yaml:"basic_auth,omitempty"`

	// Sets the `Authorization` header on every scrape request with
	// the configured bearer token. It is mutually exclusive with `bearer_token_file`.
	BearerToken common.Secret `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`

	// Sets the `Authorization` header on every scrape request with the bearer token
	// read from the configured file. It is mutually exclusive with `bearer_token`.
	BearerTokenFile string `json:"bearer_token_file,omitempty" yaml:"bearer_token_file,omitempty"`

	// Configures the scrape request's TLS settings.
	TLSConfig *common.TLSConfig `json:"tls_config,omitempty" yaml:"tls_config,omitempty"`

	// Optional proxy URL.
	ProxyURL string `json:"proxy_url,omitempty" yaml:"proxy_url,omitempty"`
}
