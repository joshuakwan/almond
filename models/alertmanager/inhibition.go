package alertmanager

import "github.com/prometheus/common/model"

// InhibitRule defines a inhibition rule
type InhibitRule struct {
	// Matchers that have to be fulfilled in the alerts to be muted.
	TargetMatch   map[string]string
	TargetMatchRe map[string]Regexp
	// Matchers for which one or more alerts have to exist for the inhibition to take effect.
	SourceMatch   map[string]string
	SourceMatchRe map[string]Regexp
	// Labels that must have an equal value in the source and target alert for the inhibition to take effect.
	Equal model.LabelNames
}
