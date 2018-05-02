package prometheus

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"sync"
)

var mu sync.Mutex

// LoadConfig loads prometheus configuration into object from a string
func LoadConfig(str string) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal([]byte(str), config)

	if err != nil {
		return nil, err
	}

	config.raw = str
	return config, nil
}

// LoadFile loads prometheus configuration into object from a file
func LoadFile(filename string) (*Config, error) {
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

func SaveConfigToFile(config *Config, filename string) error {
	mu.Lock()
	defer mu.Unlock()
	bytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

func AddScrapeConfig(configs []*ScrapeConfig, newConfig *ScrapeConfig) []*ScrapeConfig {
	return append(configs, newConfig)
}
