package alertmanager

import "github.com/prometheus/common/model"

// Receiver defines a receive who receives alerts
type Receiver struct {
	// The unique name of the receiver.
	Name string `json:"name"`

	EmailConfigs     []*EmailConfig     `json:"email_configs,omitempty" yaml:"email_configs,omitempty"`
	HipchatConfigs   []*HipchatConfig   `json:"hipchat_configs,omitempty" yaml:"hipchat_configs,omitempty"`
	PagerdutyConfigs []*PagerdutyConfig `json:"pagerduty_configs,omitempty" yaml:"pagerduty_configs,omitempty"`
	PushoverConfigs  []*PushoverConfig  `json:"pushover_configs,omitempty" yaml:"pushover_configs,omitempty"`
	SlackConfigs     []*SlackConfig     `json:"slack_configs,omitempty" yaml:"slack_configs,omitempty"`
	OpsGenieConfigs  []*OpsGenieConfig  `json:"opsgenie_configs,omitempty" yaml:"opsgenie_configs,omitempty"`
	WebhookConfigs   []*WebhookConfig   `json:"webhook_configs,omitempty" yaml:"webhook_configs,omitempty"`
	VictorOpsConfigs []*VictorOpsConfig `json:"victorops_configs,omitempty" yaml:"victorops_configs,omitempty"`
}

// EmailConfig defines the configuration of an email receiver
type EmailConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved" yaml:"send_resolved"`

	// The email address to send notifications to.
	To string `json:"to,omitempty" yaml:"to,omitempty"`
	// The sender address.
	From string `json:"from,omitempty" yaml:"from,omitempty"`
	// The SMTP host through which emails are sent.
	Smarthost string `json:"smarthost,omitempty" yaml:"smarthost,omitempty"`
	// The hostname to identify to the SMTP server.
	Hello string `json:"hello,omitempty" yaml:"hello,omitempty"`

	// SMTP authentication information.
	AuthUsername string `json:"auth_username,omitempty" yaml:"auth_username,omitempty"`
	AuthPassword string `json:"auth_password,omitempty" yaml:"auth_password,omitempty"`
	AuthSecret   Secret `json:"auth_secret,omitempty" yaml:"auth_secret,omitempty"`
	AuthIdentity string `json:"auth_identity,omitempty" yaml:"auth_identity,omitempty"`
	// The SMTP TLS requirement.
	RequireTLS bool `json:"require_tls,omitempty" yaml:"require_tls,omitempty"`
	// The HTML body of the email notification.
	HTML string `json:"html,omitempty" yaml:"html,omitempty"`
	// The text body of the email notification.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
	// Further headers email header key/value pairs. Overrides any headers
	// previously set by the notification implementation.
	Headers map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// HipchatConfig encapsulates Hipchat notification configuration
type HipchatConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved" yaml:"send_resolved"`

	// The HipChat Room ID.
	RoomID string `json:"room_id,omitempty" yaml:"room_id,omitempty"`
	// The auth token.
	AuthToken Secret `json:"auth_token,omitempty" yaml:"auth_token,omitempty"`
	// The URL to send API requests to.
	APIURL string `json:"api_url,omitempty" yaml:"api_url,omitempty"`

	// See https://www.hipchat.com/docs/apiv2/method/send_room_notification
	// A label to be shown in addition to the sender's name.
	From string `json:"from,omitempty" yaml:"from,omitempty"`
	// The message body.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	// Whether this message should trigger a user notification.
	Notify bool `json:"notify,omitempty" yaml:"notify,omitempty"`
	// Determines how the message is treated by the alertmanager and rendered inside HipChat. Valid values are 'text' and 'html'.
	MessageFormat string `json:"message_format,omitempty" yaml:"message_format,omitempty"`
	// Background color for message.
	Color string `json:"color,omitempty" yaml:"color,omitempty"`

	HTTPConfig *HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}

// PagerdutyConfig encapsulates Pagerduty notification configuration
type PagerdutyConfig struct {
	// // Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved,omitempty" yaml:"send_resolved,omitempty"`

	// The following two options are mututally exclusive.
	// The PagerDuty integration key (when using PagerDuty integration type `Events API v2`).
	RoutingKey Secret `json:"routing_key,omitempty" yaml:"routing_key,omitempty"`
	// The PagerDuty integration key (when using PagerDuty integration type `Prometheus`).
	ServiceKey Secret `json:"service_key,omitempty" yaml:"service_key,omitempty"`
	// The URL to send API requests to
	URL string `json:"url,omitempty" yaml:"url,omitempty"`
	// The client identification of the Alertmanager.
	Client string `json:"client,omitempty" yaml:"client,omitempty"`
	// A backlink to the sender of the notification.
	ClientURL string `json:"client_url,omitempty" yaml:"client_url,omitempty"`

	// A description of the incident.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Severity of the incident.
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`
	// A set of arbitrary key/value pairs that provide further detail about the incident.
	Details map[string]string `json:"details,omitempty" yaml:"details,omitempty"`

	HTTPConfig *HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}

// PushoverConfig encapsulates Pushover notification configuration
type PushoverConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved" yaml:"send_resolved"`

	// The recipient user’s user key.
	UserKey Secret `json:"user_key,omitempty" yaml:"user_key,omitempty"`
	// Your registered application’s API token, see https://pushover.net/apps
	Token Secret `json:"token,omitempty" yaml:"token,omitempty"`

	// Notification title.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// Notification message.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	// A supplementary URL shown alongside the message.
	URL string `json:"url" yaml:"url"`
	// Priority, see https://pushover.net/api#priority
	Priority string `json:"priority,omitempty" yaml:"priority,omitempty"`
	// How often the Pushover servers will send the same notification to the user.
	// Must be at least 30 seconds.
	Retry model.Duration `json:"retry,omitempty" yaml:"retry,omitempty"`
	// How long your notification will continue to be retried for, unless the user
	// acknowledges the notification.
	Expire model.Duration `json:"expire,omitempty" yaml:"expire,omitempty"`

	HTTPConfig *HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}

// SlackConfig encapsulates Slack notification configuration
type SlackConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved" yaml:"send_resolved"`

	// The Slack webhook URL.
	APIURL Secret `json:"api_url,omitempty" yaml:"api_url,omitempty"`
	// The channel or user to send notifications to.
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`

	// API request data as defined by the Slack webhook API.
	Color     string `json:"color,omitempty" yaml:"color,omitempty"`
	Username  string `json:"username,omitempty" yaml:"username,omitempty"`
	Title     string `json:"title,omitempty" yaml:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty" yaml:"title_link,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty" yaml:"icon_emoji,omitempty"`
	IconURL   string `json:"icon_url,omitempty" yaml:"icon_url,omitempty"`
	Pretext   string `json:"pretext,omitempty" yaml:"pretext,omitempty"`
	Text      string `json:"text,omitempty" yaml:"text,omitempty"`
	Fallback  string `json:"fallback,omitempty" yaml:"fallback,omitempty"`

	HTTPConfig *HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}

// OpsGenieConfig encapsulates OpsGenie notification configuration
type OpsGenieConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved" yaml:"send_resolved"`

	// The API key to use when talking to the OpsGenie API.
	APIKey Secret `json:"api_key,omitempty" yaml:"api_key,omitempty"`
	// The host to send OpsGenie API requests to.
	APIURL string `json:"api_url,omitempty" yaml:"api_url,omitempty"`

	// Alert text limited to 130 characters.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	// A description of the incident.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// A backlink to the sender of the notification.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
	// A set of arbitrary key/value pairs that provide further detail about the incident.
	Details map[string]string `json:"details,omitempty" yaml:"details,omitempty"`
	// Comma separated list of team responsible for notifications.
	Teams string `json:"teams,omitempty" yaml:"teams,omitempty"`
	// Comma separated list of tags attached to the notifications.
	Tags string `json:"tags,omitempty" yaml:"tags,omitempty"`
	// Additional alert note.
	Note string `json:"note,omitempty" yaml:"note,omitempty"`
	// Priority level of alert. Possible values are P1, P2, P3, P4, and P5.
	Priority string `json:"priority,omitempty" yaml:"priority,omitempty"`

	HTTPConfig *HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}

// WebhookConfig encapsulates Webhook notification configuration
type WebhookConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved" yaml:"send_resolved"`

	// The endpoint to send HTTP POST requests to.
	URL string `json:"url,omitempty" yaml:"url,omitempty"`

	HTTPConfig *HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}

// VictorOpsConfig encapsulates VictorOps notification configuration
type VictorOpsConfig struct {
	// Whether or not to notify about resolved alerts.
	SendResolved bool `json:"send_resolved" yaml:"send_resolved"`

	// The API key to use when talking to the VictorOps API.
	APIKey Secret `json:"api_key,omitempty" yaml:"api_key,omitempty"`
	// The VictorOps API URL.
	APIURL string `json:"api_url,omitempty" yaml:"api_url,omitempty"`

	// A key used to map the alert to a team.
	RoutingKey string `json:"routing_key,omitempty" yaml:"routing_key,omitempty"`
	// Describes the behavior of the alert (CRITICAL, WARNING, INFO).
	MessageType string `json:"message_type,omitempty" yaml:"message_type,omitempty"`
	// Contains summary of the alerted problem.
	EntityDisplayName string `json:"entity_display_name,omitempty" yaml:"entity_display_name,omitempty"`
	// Contains long explanation of the alerted problem.
	StateMessage string `json:"state_message,omitempty" yaml:"state_message,omitempty"`
	// The monitoring tool the state message is from.
	MonitoringTool string `json:"monitoring_tool,omitempty" yaml:"monitoring_tool,omitempty"`

	HTTPConfig *HTTPConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
}
