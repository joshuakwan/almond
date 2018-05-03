package models

type GrafanaDashboard struct {
	UID  string `json:"uid,omitempty"`
	URL  string `json:"url,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type Service struct {
	ServiceID   string            `json:"service_id,omitempty"`   // consul service id
	ServiceName string            `json:"service_name,omitempty"` // consul service name
	Dashboard   *GrafanaDashboard `json:"dashboard,omitempty"`    // grafana dashboard info
}

type Tenant struct {
	Name                   string     `json:"name"`
	Description            string     `json:"description,omitempty"`
	Services               []*Service `json:"services,omitempty"`
	GrafanaOrgName         string     `json:"grafana_org,omitempty"` // grafana org info
	GrafanaOrgID           int        `json:"grafana_org_id"`
	GrafanaOrgUser         string     `json:"grafana_org_user"` // admin user name for the tenant
	GrafanaOrgUserID       int        `json:"grafana_org_user_id"`
	GrafanaOrgUserPassword string     `json:"grafana_org_user_password"` // give it to the tenant
	GrafanaOrgAdminKey     string     `json:"grafana_org_admin_key"`     // the api key of the admin
	GrafanaURL             string     `json:"grafana_url,omitempty"`     // target grafana url
}

type ServiceRegistration struct {
	Port    int    `json:"port"`
	Address string `json:"address"`
	Type    string `json:"type"`
}
