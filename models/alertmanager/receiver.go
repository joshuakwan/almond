package alertmanager

import "github.com/prometheus/common/model"

// Receiver defines a receive who receives alerts
type Receiver struct {
	// The unique name of the receiver.
	Name string

	EmailConfigs     []*EmailConfig
	HipchatConfigs   []*HipchatConfig
	PagerdutyConfigs []*PagerdutyConfig
	PushoverConfigs  []*PushoverConfig
	SlackConfigs     []*SlackConfig
	OpsGenieConfigs  []*OpsGenieConfig
	WebhookConfigs   []*WebhookConfig
	VictorOpsConfigs []*VictorOpsConfig
}

// EmailConfig defines the configuration of an email receiver
type EmailConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool

	// The email address to send notifications to.
	To string
	// The sender address.
	From string
	// The SMTP host through which emails are sent.
	Smarthost string
	// The hostname to identify to the SMTP server.
	Hello string

	// SMTP authentication information.
	AuthUsername string
	AuthPassword string
	AuthSecret   Secret
	AuthIdentity string
	// The SMTP TLS requirement.
	RequireTLS bool
	// The HTML body of the email notification.
	HTML string
	// The text body of the email notification.
	Text string
	// Further headers email header key/value pairs. Overrides any headers
	// previously set by the notification implementation.
	Headers map[string]string
}

// HipchatConfig encapsulates Hipchat notification configuration
type HipchatConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool

	// The HipChat Room ID.
	RoomID string
	// The auth token.
	AuthToken Secret
	// The URL to send API requests to.
	APIURL string

	// See https://www.hipchat.com/docs/apiv2/method/send_room_notification
	// A label to be shown in addition to the sender's name.
	From string
	// The message body.
	Message string
	// Whether this message should trigger a user notification.
	Notify bool
	// Determines how the message is treated by the alertmanager and rendered inside HipChat. Valid values are 'text' and 'html'.
	MessageFormat string
	// Background color for message.
	Color string

	HTTPConfig *HTTPConfig
}

// PagerdutyConfig encapsulates Pagerduty notification configuration
type PagerdutyConfig struct {
	// // Whether or not to notify about resolved alerts.
	SendResolved bool

	// The following two options are mututally exclusive.
	// The PagerDuty integration key (when using PagerDuty integration type `Events API v2`).
	RoutingKey Secret
	// The PagerDuty integration key (when using PagerDuty integration type `Prometheus`).
	ServiceKey Secret
	// The URL to send API requests to
	URL string
	// The client identification of the Alertmanager.
	Client string
	// A backlink to the sender of the notification.
	ClientURL string

	// A description of the incident.
	Description string
	// Severity of the incident.
	Severity string
	// A set of arbitrary key/value pairs that provide further detail about the incident.
	Details map[string]string

	HTTPConfig *HTTPConfig
}

// PushoverConfig encapsulates Pushover notification configuration
type PushoverConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool

	// The recipient user’s user key.
	UserKey Secret
	// Your registered application’s API token, see https://pushover.net/apps
	Token Secret

	// Notification title.
	Title string
	// Notification message.
	Message string
	// A supplementary URL shown alongside the message.
	URL string
	// Priority, see https://pushover.net/api#priority
	Priority string
	// How often the Pushover servers will send the same notification to the user.
	// Must be at least 30 seconds.
	Retry model.Duration
	// How long your notification will continue to be retried for, unless the user
	// acknowledges the notification.
	Expire model.Duration

	HTTPConfig *HTTPConfig
}

// SlackConfig encapsulates Slack notification configuration
type SlackConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool

	// The Slack webhook URL.
	APIURL Secret
	// The channel or user to send notifications to.
	Channel string

	// API request data as defined by the Slack webhook API.
	Color     string
	Username  string
	Title     string
	TitleLink string
	IconEmoji string
	IconURL   string
	Pretext   string
	Text      string
	Fallback  string

	HTTPConfig *HTTPConfig
}

// OpsGenieConfig encapsulates OpsGenie notification configuration
type OpsGenieConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool

	// The API key to use when talking to the OpsGenie API.
	APIKey Secret
	// The host to send OpsGenie API requests to.
	APIURL string

	// Alert text limited to 130 characters.
	Message string
	// A description of the incident.
	Description string
	// A backlink to the sender of the notification.
	Source string
	// A set of arbitrary key/value pairs that provide further detail about the incident.
	Details map[string]string
	// Comma separated list of team responsible for notifications.
	Teams string
	// Comma separated list of tags attached to the notifications.
	Tags string
	// Additional alert note.
	Note string
	// Priority level of alert. Possible values are P1, P2, P3, P4, and P5.
	Priority string

	HTTPConfig *HTTPConfig
}

// WebhookConfig encapsulates Webhook notification configuration
type WebhookConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool

	// The endpoint to send HTTP POST requests to.
	URL string

	HTTPConfig *HTTPConfig
}

// VictorOpsConfig encapsulates VictorOps notification configuration
type VictorOpsConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool

	// The API key to use when talking to the VictorOps API.
	APIKey Secret
	// The VictorOps API URL.
	APIURL string

	// A key used to map the alert to a team.
	RoutingKey string
	// Describes the behavior of the alert (CRITICAL, WARNING, INFO).
	MessageType string
	// Contains summary of the alerted problem.
	EntityDisplayName string
	// Contains long explanation of the alerted problem.
	StateMessage string
	// The monitoring tool the state message is from.
	MonitoringTool string

	HTTPConfig *HTTPConfig
}
