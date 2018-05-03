package controllers

import (
	grafana_api "github.com/joshuakwan/grafana-client/api"
)

func getGrafanaOrganizationClients() map[int]*grafana_api.Client {
	var clients = make(map[int]*grafana_api.Client)
	tenants, err := getTenants()
	if err != nil {
		panic(err)
	}
	for _, tenant := range (tenants) {
		clients[tenant.GrafanaOrgID] = &grafana_api.Client{
			GrafanaURL:    tenant.GrafanaURL,
			BearerToken:   tenant.GrafanaOrgAdminKey,
			AdminUser:     "",
			AdminPassword: "",
		}
	}

	return clients
}

func getGrafanaClient(url string, token string, adminUser string, adminPassword string) *grafana_api.Client {
	return &grafana_api.Client{
		GrafanaURL:    url,
		BearerToken:   token,
		AdminUser:     adminUser,
		AdminPassword: adminPassword,
	}
}
