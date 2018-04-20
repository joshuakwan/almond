package alertmanager

import (
	"io/ioutil"
	"github.com/prometheus/alertmanager/config"
	"gopkg.in/yaml.v2"
	"sync"
	"github.com/getlantern/deepcopy"
	"github.com/fatih/structs"
	"strings"
	"errors"
)

var mu sync.Mutex

func SaveConfigToFile(config *config.Config, filename string) error {
	mu.Lock()
	defer mu.Unlock()
	bytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

func Update(dst *config.GlobalConfig,src *config.GlobalConfig) {
	deepcopy.Copy(dst, src)
}

func Delete(g *config.GlobalConfig,key string) {
	for _, field := range (structs.Fields(g)) {
		tagName := field.Tag("json")
		parts := strings.Split(tagName, ",")
		if key == parts[0] {
			field.Zero()
			break
		}
	}
}

func AddInhibitRule(rules []*config.InhibitRule, newRule *config.InhibitRule) []*config.InhibitRule {
	return append(rules, newRule)
}

func RemoveInhibitRule(rules []*config.InhibitRule, index int) ([]*config.InhibitRule, error) {
	if index >= len(rules) || index < 0 {
		return nil, errors.New("Index " + string(index) + " not in the right range")
	}

	copy(rules[index:], rules[index+1:])
	return rules[:len(rules)-1], nil
}

func AddReceiver(receivers []*config.Receiver, newReceiver *config.Receiver) ([]*config.Receiver, error) {
	for _, receiver := range (receivers) {
		if receiver.Name == newReceiver.Name {
			return nil, errors.New("Receiver " + newReceiver.Name + " already exists")
		}
	}

	return append(receivers, newReceiver), nil
}

func RemoveReceiver(receivers []*config.Receiver, name string) ([]*config.Receiver, error) {
	for i, receiver := range (receivers) {
		if receiver.Name == name {
			copy(receivers[i:], receivers[i+1:])
			return receivers[:len(receivers)-1], nil
		}
	}

	return nil, errors.New("Receiver " + name + " not found")
}

func AddSubroute(route *config.Route, subroute *config.Route) *config.Route {
	route.Routes = append(route.Routes, subroute)
	return route
}

func RemoveSubroute(route *config.Route, index int) (*config.Route, error) {
	if index >= len(route.Routes) || index < 0 {
		return nil, errors.New("Index " + string(index) + " not in the right range")
	}
	subroutes := route.Routes
	copy(subroutes[index:], subroutes[index+1:])
	route.Routes = subroutes[:len(subroutes)-1]
	return route, nil
}
