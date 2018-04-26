package facade

import (
	"github.com/joshuakwan/almond/models/almond"
	consul_api "github.com/hashicorp/consul/api"
	grafana_api "github.com/joshuakwan/grafana-client/api"
	"log"
	"encoding/json"
)

func getConsulClient(url string) *consul_api.Client {
	log.Println("Initializing consul API client")
	client, err := consul_api.NewClient(&consul_api.Config{Address: url})
	if err != nil {
		panic(err)
	}
	log.Println(client)
	return client
}

func putTenant(consul *consul_api.Client,tenant *almond.Tenant) error {
	data, err := json.Marshal(tenant)
	if err != nil {
		return err
	}

	p := &consul_api.KVPair{Key: tenantsRoot + tenant.Name, Value: data}
	_, err = consul.KV().Put(p, nil)
	return err
}

func getGrafanClient(url string, token string, adminUser string, adminPassword string) *grafana_api.Client {
	return &grafana_api.Client{
		GrafanaURL:    url,
		BearerToken:   token,
		AdminUser:     adminUser,
		AdminPassword: adminPassword,
	}
}
