package prometheus

import "github.com/prometheus/common/model"

// Global represents global configuration of prometheus
type Global struct {
	// How frequently to scrape targets by default.
	// default = 1m
	ScrapeInterval model.Duration `json:"scrape_interval,omitempty" yaml:"scrape_interval,omitempty"`

	// How long until a scrape request times out.
	// default = 10s
	ScrapeTimeout model.Duration `json:"scrape_timeout,omitempty" yaml:"scrape_timeout,omitempty"`

	// How frequently to evaluate rules.
	// default = 1m
	EvaluationInterval model.Duration `json:"evaluation_interval,omitempty" yaml:"evaluation_interval,omitempty"`

	// The labels to add to any time series or alerts when communicating with
	// external systems (federation, remote storage, Alertmanager).
	ExternalLabels map[string]string `json:"external_labels,omitempty" yaml:"external_labels,omitempty"`
}
