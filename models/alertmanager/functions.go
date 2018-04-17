package alertmanager

import (
	"io/ioutil"
	"regexp"
	"time"

	"github.com/prometheus/common/model"
	"gopkg.in/yaml.v2"
)

var defaultGlobal = Global{
	ResolveTimeout:  model.Duration(5 * time.Minute),
	SMTPHello:       "localhost",
	SMTPRequireTLS:  true,
	VictorOpsAPIURL: "https://alert.victorops.com/integrations/generic/20131114/alert/",
	PagerdutyURL:    "https://events.pagerduty.com/v2/enqueue",
	OpsGenieAPIURL:  "https://api.opsgenie.com/",
	HipchatAPIURL:   "https://api.hipchat.com/",
}

var defaultRoute = Route{
	Continue:       false,
	GroupWait:      model.Duration(30 * time.Second),
	GroupInterval:  model.Duration(5 * time.Minute),
	RepeatInterval: model.Duration(4 * time.Hour),
}

var defaultEmailConfig = EmailConfig{
	SendResolved: false,
	HTML:         `{{ template "email.default.html" . }}`,
}

var defaultHipchatConfig = HipchatConfig{
	SendResolved:  false,
	From:          `{{ template "hipchat.default.from" . }}`,
	Message:       `{{ template "hipchat.default.message" . }}`,
	Notify:        false,
	MessageFormat: "text",
	Color:         `{{ if eq .Status "firing" }}red{{ else }}green{{ end }}`,
}

var defaultPagerdutyConfig = PagerdutyConfig{
	SendResolved: true,
	Client:       `{{ template "pagerduty.default.client" . }}`,
	ClientURL:    `{{ template "pagerduty.default.clientURL" . }}`,
	Description:  `{{ template "pagerduty.default.description" .}}`,
	Severity:     "error",
	Details: map[string]string{
		"firing":       `{{ template "pagerduty.default.instances" .Alerts.Firing }}`,
		"resolved":     `{{ template "pagerduty.default.instances" .Alerts.Resolved }}`,
		"num_firing":   `{{ .Alerts.Firing | len }}`,
		"num_resolved": `{{ .Alerts.Resolved | len }}`,
	},
}

var defaultPushoverConfig = PushoverConfig{
	SendResolved: true,
	Title:        `{{ template "pushover.default.title" . }}`,
	Message:      `{{ template "pushover.default.message" . }}`,
	URL:          `{{ template "pushover.default.url" . }}`,
	Priority:     `{{ if eq .Status "firing" }}2{{ else }}0{{ end }}`,
	Retry:        model.Duration(1 * time.Minute),
	Expire:       model.Duration(1 * time.Hour),
}

var defaultSlackConfig = SlackConfig{
	SendResolved: false,
	Color:        `{{ if eq .Status "firing" }}danger{{ else }}good{{ end }}`,
	Username:     `{{ template "slack.default.username" . }}`,
	Title:        `{{ template "slack.default.title" . }}`,
	TitleLink:    `{{ template "slack.default.titlelink" . }}`,
	Pretext:      `{{ template "slack.default.pretext" . }}`,
	Text:         `{{ template "slack.default.text" . }}`,
	Fallback:     `{{ template "slack.default.fallback" . }}`,
}

var defaultOpsGenieConfig = OpsGenieConfig{
	SendResolved: true,
	Description:  `{{ template "opsgenie.default.description" . }}`,
	Source:       `{{ template "opsgenie.default.source" . }}`,
}

var defaultVictorOpsConfig = VictorOpsConfig{
	SendResolved:      true,
	MessageType:       "CRITICAL",
	EntityDisplayName: `{{ template "victorops.default.entity_display_name" . }}`,
	StateMessage:      `{{ template "victorops.default.state_message" . }}`,
	MonitoringTool:    `{{ template "victorops.default.monitoring_tool" . }}`,
}

var defaultWebhookConfig = WebhookConfig{
	SendResolved: true,
}

func (re *Regexp) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	regex, err := regexp.Compile("^(?:" + s + ")$")
	if err != nil {
		return err
	}
	re.Regexp = regex
	return nil
}

func (re Regexp) MarshalYAML() (interface{}, error) {
	if re.Regexp != nil {
		return re.String(), nil
	}
	return nil, nil
}

func LoadConfig(str string) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal([]byte(str), config)

	if err != nil {
		return nil, err
	}

	config.raw = str
	return config, nil
}

func LoadConfigFromFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config, err := LoadConfig(string(content))
	if err != nil {
		return nil, err
	}
	return config, err
}
