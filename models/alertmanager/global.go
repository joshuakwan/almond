package alertmanager

import (
	"github.com/joshuakwan/almond/models/common"
	"github.com/prometheus/common/model"
)

// Global defines the global part of alertmanager configuration
type Global struct {
	// ResolveTimeout is the time after which an alert is declared resolved
	// if it has not been updated.
	ResolveTimeout model.Duration `json:"resolve_timeout,omitempty" yaml:"resolve_timeout,omitempty"`

	// The default SMTP From header field.
	SMTPFrom string `json:"smtp_from,omitempty" yaml:"smtp_from,omitempty"`
	// The default SMTP smarthost used for sending emails, including port number.
	// Port number usually is 25, or 587 for SMTP over TLS (sometimes referred to as STARTTLS).
	// Example: smtp.example.org:587
	SMTPSmartHost string `json:"smtp_smarthost,omitempty" yaml:"smtp_smarthost,omitempty"`
	// The default hostname to identify to the SMTP server.
	SMTPHello string `json:"smtp_hello,omitempty" yaml:"smtp_hello,omitempty"`
	// SMTP Auth using LOGIN and PLAIN.
	SMTPAuthUsername string `json:"smtp_auth_username,omitempty" yaml:"smtp_auth_username,omitempty"`
	SMTPAuthPassword string `json:"smtp_auth_password,omitempty" yaml:"smtp_auth_password,omitempty"`
	// SMTP Auth using PLAIN.
	SMTPAuthIdentity string `json:"smtp_auth_identity,omitempty" yaml:"smtp_auth_identity,omitempty"`
	// SMTP Auth using CRAM-MD5.
	SMTPAuthSecret common.Secret `json:"smtp_auth_secret,omitempty" yaml:"smtp_auth_secret,omitempty"`
	// The default SMTP TLS requirement.
	SMTPRequireTLS bool `json:"smtp_require_tls,omitempty" yaml:"smtp_require_tls,omitempty"`

	SlackAPIURL      string        `json:"slack_api_url,omitempty" yaml:"slack_api_url,omitempty"`
	VictorOpsAPIKey  string        `json:"victorops_api_key,omitempty" yaml:"victorops_api_key,omitempty"`
	VictorOpsAPIURL  string        `json:"victorops_api_url,omitempty" yaml:"victorops_api_url,omitempty"`
	PagerdutyURL     string        `json:"pagerduty_url,omitempty" yaml:"pagerduty_url,omitempty"`
	OpsGenieAPIKey   string        `json:"opsgenie_api_key,omitempty" yaml:"opsgenie_api_key,omitempty"`
	OpsGenieAPIURL   string        `json:"opsgenie_api_url,omitempty" yaml:"opsgenie_api_url,omitempty"`
	HipchatAPIURL    string        `json:"hipchat_api_url,omitempty" yaml:"hipchat_api_url,omitempty"`
	HipchatAuthToken common.Secret `json:"hipchat_auth_token,omitempty" yaml:"hipchat_auth_token,omitempty"`

	HTTPConfig HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}
