package facade

import (
	"github.com/joshuakwan/almond/models/almond"
	consul_api "github.com/hashicorp/consul/api"
	grafana_api "github.com/joshuakwan/grafana-client/api"
	"log"
	"encoding/json"
	"errors"
	"fmt"
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

func getConsulKVData(consul *consul_api.Client, key string) ([]byte, error) {
	pair, _, err := consul.KV().Get(key, nil)
	if pair == nil {
		return nil, errors.New(fmt.Sprint("%v not found", key))
	}
	return pair.Value, err
}

func listConsulKVData(consul *consul_api.Client, key string) ([][]byte, error) {
	pairs, _, err := consul.KV().List(key, nil)
	if pairs == nil {
		return nil, errors.New(fmt.Sprint("%v not found", key))
	}
	// need to discard pairs[0]
	var data [][]byte

	if len(pairs) == 1 {
		return data, nil
	}

	for _, pair := range (pairs[1:]) {
		data = append(data, pair.Value)
	}

	return data, err
}

func getTenant(tenantName string) (*almond.Tenant, error) {
	data, err := getConsulKVData(consulClient, tenantsRoot+tenantName)
	if err != nil {
		return nil, err
	}
	var tenant almond.Tenant
	err = json.Unmarshal(data, &tenant)
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func getTenants() ([]*almond.Tenant, error) {
	data, err := listConsulKVData(consulClient, tenantsRoot)
	if err != nil {
		return nil, err
	}

	var tenants []*almond.Tenant
	for _, d := range (data) {
		var tenant almond.Tenant
		err = json.Unmarshal(d, &tenant)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, &tenant)
	}

	return tenants, nil
}

func putTenant(consul *consul_api.Client, tenant *almond.Tenant) error {
	data, err := json.Marshal(tenant)
	if err != nil {
		return err
	}

	p := &consul_api.KVPair{Key: tenantsRoot + tenant.Name, Value: data}
	_, err = consul.KV().Put(p, nil)
	return err
}

func checkIfTenantExists(tenantName string) bool {
	tenant, _ := getTenant(tenantName)

	if tenant == nil {
		return false
	} else {
		return true
	}
}

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
