package prometheus

import "io/ioutil"
import "gopkg.in/yaml.v2"

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

// LoadConfigFromFile loads prometheus configuration into object from a file
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
