package common

import (
	"regexp"
	"github.com/getlantern/deepcopy"
)

// Secret represents a Secret string
type Secret string

// Regexp represents a Regexp
type Regexp struct {
	*regexp.Regexp
}

// BasicAuth defines Basic Auth information
type BasicAuth struct {
	Username string `json:"username" yaml:"username"`
	Password Secret `json:"password" yaml:"username"`
}

// TLSConfig defines TLS configuration items
type TLSConfig struct {
	// CA certificate to validate the server certificate with.
	CAFile string `json:"ca_file" yaml:"ca_file"`
	// Certificate and key files for client cert authentication to the server.
	CertFile string `json:"cert_file" yaml:"cert_file"`
	KeyFile  string `json:"key_file" yaml:"key_file"`
	// ServerName extension to indicate the name of the server.
	ServerName string `json:"server_name" yaml:"server_name"`
	// Disable validation of the server certificate.
	InsecureSkipVerity bool `json:"insecure_skip_verity" yaml:"insecure_skip_verity"`
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

func Update(dst interface{}, src interface{}) {
	deepcopy.Copy(dst, src)
}