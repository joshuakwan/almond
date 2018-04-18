package prometheus

import "github.com/joshuakwan/almond/models/common"

// Config encapsulates Prometheus configuration
type Config struct {
	Global        *Global         `json:"global,omitempty" yaml:"global,omitempty"`
	RuleFiles     []string        `json:"rule_files,omitempty" yaml:"rule_files,omitempty"`
	ScrapeConfigs []*ScrapeConfig `json:"scrape_configs,omitempty" yaml:"scrape_configs,omitempty"`
	Alerting      *Alerting       `json:"alerting,omitempty" yaml:"alerting,omitempty"`
	RemoteWrite   *RemoteWrite    `json:"remote_write,omitempty" yaml:"remote_write,omitempty"`
	RemoteRead    *RemoteRead     `json:"remote_read,omitempty" yaml:"remote_read,omitempty"`

	raw string
}

// StaticConfig represents static_configs
type StaticConfig struct {
	// The targets specified by the static config.
	Targets []string `json:"targets,omitempty" yaml:"targets,omitempty"`

	// Labels assigned to all metrics scraped from the targets.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// RelabelConfig represents relabel_configs
type RelabelConfig struct {
	// The source labels select values from existing labels. Their content is concatenated
	// using the configured separator and matched against the configured regular expression
	// for the replace, keep, and drop actions.
	SourceLabels string `json:"source_labels,omitempty" yaml:"source_labels,omitempty"`

	// Separator placed between concatenated source label values.
	Separator string `json:"separator,omitempty" yaml:"separator,omitempty"` // default = ;

	// Label to which the resulting value is written in a replace action.
	// It is mandatory for replace actions. Regex capture groups are available.
	TargetLabel string `json:"target_label,omitempty" yaml:"target_label,omitempty"`

	// Regular expression against which the extracted value is matched.
	Regex common.Regexp `json:"regex,omitempty" yaml:"regex,omitempty"` // default = (.*)

	// Modulus to take of the hash of the source label values.
	Modulus uint64 `json:"modulus,omitempty" yaml:"modulus,omitempty"`

	// Replacement value against which a regex replace is performed if the
	// regular expression matches. Regex capture groups are available.
	Replacement string `json:"replacement,omitempty" yaml:"replacement,omitempty"` // default = $1

	// Action to perform based on regex matching.
	Action string `json:"action,omitempty" yaml:"action,omitempty"` // default = replace
}
