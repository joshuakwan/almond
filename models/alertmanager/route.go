package alertmanager

import (
	"github.com/prometheus/common/model"
	"errors"
)

// Route represents a routing rule
type Route struct {
	// Should match the receiver's name
	Receiver string `json:"receiver,omitempty" yaml:"receiver,omitempty"`

	GroupBy []model.LabelName `json:"group_by,omitempty" yaml:"group_by,omitempty"`
	// Whether an alert should continue matching subsequent sibling nodes.
	Continue bool `json:"continue,omitempty" yaml:"continue,omitempty"`

	// A set of equality matchers an alert has to fulfill to match the node.
	Match map[string]string `json:"match,omitempty" yaml:"match,omitempty"`
	// A set of regex-matchers an alert has to fulfill to match the node.
	MatchRe map[string]string `json:"match_re,omitempty" yaml:"match_re,omitempty"`

	// How long to initially wait to send a notification for a group
	// of alerts. Allows to wait for an inhibiting alert to arrive or collect
	// more initial alerts for the same group. (Usually ~0s to few minutes.)
	GroupWait model.Duration `json:"group_wait,omitempty" yaml:"group_wait,omitempty"`
	// How long to wait before sending a notification about new alerts that
	// are added to a group of alerts for which an initial notification has
	// already been sent. (Usually ~5m or more.)
	GroupInterval model.Duration `json:"group_interval,omitempty" yaml:"group_interval,omitempty"`
	// How long to wait before sending a notification again if it has already
	// been sent successfully for an alert. (Usually ~3h or more).
	RepeatInterval model.Duration `json:"repeat_interval,omitempty" yaml:"repeat_interval,omitempty"`

	// 0 or more child routes
	Routes []*Route `json:"routes,omitempty" yaml:"routes,omitempty"`
}

func Addubroute(route *Route, subroute *Route) *Route {
	route.Routes = append(route.Routes, subroute)
	return route
}

func RemoveSubroute(route *Route, index int) (*Route, error) {
	if index >= len(route.Routes) || index < 0 {
		return nil, errors.New("Index " + string(index) + " not in the right range")
	}
	subroutes := route.Routes
	copy(subroutes[index:], subroutes[index+1:])
	route.Routes = subroutes[:len(subroutes)-1]
	return route, nil
}
