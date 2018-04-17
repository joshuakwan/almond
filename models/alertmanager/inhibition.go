package alertmanager

import "github.com/prometheus/common/model"

// InhibitRule defines a inhibition rule
type InhibitRule struct {
	// Matchers that have to be fulfilled in the alerts to be muted.
	TargetMatch   map[string]string `json:"target_match,omitempty" yaml:"target_match,omitempty"`
	TargetMatchRe map[string]Regexp `json:"target_match_re,omitempty" yaml:"target_match_re,omitempty"`
	// Matchers for which one or more alerts have to exist for the inhibition to take effect.
	SourceMatch   map[string]string `json:"source_match,omitempty" yaml:"source_match,omitempty"`
	SourceMatchRe map[string]Regexp `json:"source_match_re,omitempty" yaml:"source_match_re,omitempty"`
	// Labels that must have an equal value in the source and target alert for the inhibition to take effect.
	Equal model.LabelNames `json:"equal,omitempty" yaml:"equal,omitempty"`
}
