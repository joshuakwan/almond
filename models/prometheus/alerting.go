package prometheus

import (
	"github.com/joshuakwan/almond/models/common"
	"github.com/prometheus/common/model"
)

// Alerting represents alerting section
type Alerting struct {
	AlertRelabelConfigs []*RelabelConfig
	AlertManagers       []*AlertManagerConfig
}

// AlertManagerConfig represents alerting.alertmanagers section
type AlertManagerConfig struct {
	// Per-target Alertmanager timeout when pushing alerts.
	Timeout model.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"` //default = 10s

	// Prefix for the HTTP path alerts are pushed to.
	PathPrefix string `json:"path_prefix,omitempty" yaml:"path_prefix,omitempty"` //default = /

	// Configures the protocol scheme used for requests.
	// [ scheme: <scheme> | default = http ]
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`

	// Sets the `Authorization` header on every request with the
	// configured username and password.
	BasicAuth *common.BasicAuth `json:"basic_auth,omitempty" yaml:"basic_auth,omitempty"`

	// Sets the `Authorization` header on every request with
	// the configured bearer token. It is mutually exclusive with `bearer_token_file`.
	BearerToken common.Secret `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`

	// Sets the `Authorization` header on every request with the bearer token
	// read from the configured file. It is mutually exclusive with `bearer_token`.
	BearerTokenFile string `json:"bearer_token_file,omitempty" yaml:"bearer_token_file,omitempty"`

	// Configures the scrape request's TLS settings.
	TLSConfig *common.TLSConfig `json:"tls_config,omitempty" yaml:"tls_config,omitempty"`

	// Optional proxy URL.
	ProxyURL string `json:"proxy_url,omitempty" yaml:"proxy_url,omitempty"`

	// List of Azure service discovery configurations.
	AzureSdConfigs []*AzureSdConfig `json:"azure_sd_configs,omitempty" yaml:"azure_sd_configs,omitempty"`

	// List of Consul service discovery configurations.
	ConsulSdConfigs []*ConsulSdConfig `json:"consul_sd_configs,omitempty" yaml:"consul_sd_configs,omitempty"`

	// List of DNS service discovery configurations.
	DNSSdConfigs []*DNSSdConfig `json:"dns_sd_configs,omitempty" yaml:"dns_sd_configs,omitempty"`

	// List of EC2 service discovery configurations.
	EC2SdConfigs []*EC2SdConfig `json:"ec2_sd_configs,omitempty" yaml:"ec2_sd_configs,omitempty"`

	// List of OpenStack service discovery configurations.
	OpenStackSdConfigs []*OpenStackSdConfig `json:"openstack_sd_configs,omitempty" yaml:"openstack_sd_configs,omitempty"`

	// List of file service discovery configurations.
	FileSdConfigs []*FileSdConfig `json:"file_sd_configs,omitempty" yaml:"file_sd_configs,omitempty"`

	// List of GCE service discovery configurations.
	GCESdConfigs []*GCESdConfig `json:"gce_sd_configs,omitempty" yaml:"gce_sd_configs,omitempty"`

	// List of Kubernetes service discovery configurations.
	KubernetesSdConfigs []*KubernetesSdConfig `json:"kubernetes_sd_configs,omitempty" yaml:"kubernetes_sd_configs,omitempty"`

	// List of Marathon service discovery configurations.
	MarathonSdConfigs []*MarathonSdConfig `json:"marathon_sd_configs,omitempty" yaml:"marathon_sd_configs,omitempty"`

	// List of AirBnB's Nerve service discovery configurations.
	NerveSdConfigs []*NerveSdConfig `json:"nerve_sd_configs,omitempty" yaml:"nerve_sd_configs,omitempty"`

	// List of Zookeeper Serverset service discovery configurations.
	ServersetSdConfigs []*ServersetSdConfig `json:"serverset_sd_configs,omitempty" yaml:"serverset_sd_configs,omitempty"`

	// List of Triton service discovery configurations.
	TritonSdConfigs []*TritonSdConfig `json:"triton_sd_configs,omitempty" yaml:"triton_sd_configs,omitempty"`

	// List of labeled statically configured targets for this job.
	StaticConfigs []*StaticConfig `json:"static_configs,omitempty" yaml:"static_configs,omitempty"`

	// List of target relabel configurations.
	RelabelConfigs []*RelabelConfig `json:"relabel_configs,omitempty" yaml:"relabel_configs,omitempty"`
}
