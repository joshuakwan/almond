package alertmanager

import "github.com/prometheus/common/model"

// Route represents a routing rule
type Route struct {
	// Should match the receiver's name
	Receiver string
	GroupBy  []model.LabelName
	// Whether an alert should continue matching subsequent sibling nodes.
	Continue bool

	// A set of equality matchers an alert has to fulfill to match the node.
	Match map[string]string
	// A set of regex-matchers an alert has to fulfill to match the node.
	MatchRe map[string]Regexp

	// How long to initially wait to send a notification for a group
	// of alerts. Allows to wait for an inhibiting alert to arrive or collect
	// more initial alerts for the same group. (Usually ~0s to few minutes.)
	GroupWait *model.Duration
	// How long to wait before sending a notification about new alerts that
	// are added to a group of alerts for which an initial notification has
	// already been sent. (Usually ~5m or more.)
	GroupInterval *model.Duration
	// How long to wait before sending a notification again if it has already
	// been sent successfully for an alert. (Usually ~3h or more).
	RepeatInterval *model.Duration

	// 0 or more child routes
	Routes []*Route
}
