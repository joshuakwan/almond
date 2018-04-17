package alertmanager

import "github.com/prometheus/common/model"

// Global defines the global part of alertmanager configuration
type Global struct {
	// ResolveTimeout is the time after which an alert is declared resolved
	// if it has not been updated.
	ResolveTimeout model.Duration

	// The default SMTP From header field.
	SMTPFrom string
	// The default SMTP smarthost used for sending emails, including port number.
	// Port number usually is 25, or 587 for SMTP over TLS (sometimes referred to as STARTTLS).
	// Example: smtp.example.org:587
	SMTPSmartHost string
	// The default hostname to identify to the SMTP server.
	SMTPHello string
	// SMTP Auth using LOGIN and PLAIN.
	SMTPAuthUsername string
	SMTPAuthPassword string
	// SMTP Auth using PLAIN.
	SMTPAuthIdentity string
	// SMTP Auth using CRAM-MD5.
	SMTPAuthSecret Secret
	// The default SMTP TLS requirement.
	SMTPRequireTLS bool

	SlackAPIURL      string
	VictorOpsAPIKey  string
	VictorOpsAPIURL  string
	PagerdutyURL     string
	OpsGenieAPIKey   string
	OpsGenieAPIURL   string
	HipchatAPIURL    string
	HipchatAuthToken Secret
	HTTPConfig       HTTPConfig
}
