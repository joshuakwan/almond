package controllers

import (
	"github.com/joshuakwan/almond/models"
	consul_api "github.com/hashicorp/consul/api"
	"log"
	"errors"
	"fmt"
)

// RegisterDashboard register a dashboard (its json data) to consul under grafana_dashboards/
func RegisterDashboard(name string, jsondata []byte) error {
	key := dashboardRoot + name
	log.Println("register dashboard at", key)

	p := &consul_api.KVPair{Key: key, Value: jsondata}
	_, err := consulClient.KV().Put(p, nil)
	return err
}

// things to do:
//   1. check if the tenant already exists
//   2. create a new grafana org
//   3. TODO generate an admin key of the new org - issue #7
//   4. create a new global user
//   5. assign the user to the new org
//   6. create datasource
//   7. create a new folder in consul kv store tenants/{tenant name}
func CreateTenant(tenant *models.Tenant) (*models.Tenant, error) {
	// 1. check if the tenant already exists
	log.Println("check if tenant " + tenant.Name + " exists")
	if checkIfTenantExists(tenant.Name) == true {
		return nil, errors.New(fmt.Sprintf("tenant %v already exists", tenant.Name))
	}

	log.Println("create tenant", tenant.Name)
	doer := &commander{}

	// 2. create a new grafana org
	cmdGrafanaOrgCreation := &grafanaOrgCreationCommand{
		grafana:    grafanaClient,
		grafanaURL: grafanaUrl,
		tenant:     tenant,
	}
	doer.addCommand(cmdGrafanaOrgCreation)

	// 3. generate admin key of the new org
	cmdGrafanaAdminKeyCreation := &grafanaAdminKeyCreationCommand{
		grafana: grafanaClient,
		tenant:  tenant,
	}
	doer.addCommand(cmdGrafanaAdminKeyCreation)

	// 4. create a new org user
	cmdGrafanaUserCreation := &grafanaUserCreationCommand{
		grafana: grafanaClient,
		tenant:  tenant,
	}
	doer.addCommand(cmdGrafanaUserCreation)

	// 5. assign the user to the new org
	cmdGrafanaUserAssignment := &grafanaUserAssignmentCommand{
		grafana: grafanaClient,
		tenant:  tenant,
	}
	doer.addCommand(cmdGrafanaUserAssignment)

	// 6. create a new tenant entry in consul kv store as tenants/{tenant name}
	cmdTenantPut := &putTenantCommand{
		consul: consulClient,
		tenant: tenant,
	}
	doer.addCommand(cmdTenantPut)

	// 7. add a prometheus scrape config
	cmdPrometheusScrapeConfig := &prometheusScrapeConfigCreationCommand{
		tenant: tenant,
	}
	doer.addCommand(cmdPrometheusScrapeConfig)

	// 8. create datasource
	cmdGrafanaDatasourceCreation := &grafanaDatasourceCreationCommand{
		grafanaClients: grafanaOrgClients,
		tenant:         tenant,
	}
	doer.addCommand(cmdGrafanaDatasourceCreation)

	err := doer.execute()
	if err != nil {
		doer.rollback()
		return nil, err
	}

	return tenant, nil
}

// RegisterService register a service to a tenant
// 3 parts:
//   1. register the service to consul
//   2. do grafana stuff
//   3. link the service id to the tenant in the kv store
// TODO happy path done, need to make it robust and consistent
func RegisterService(tenantName string, service *models.ServiceRegistration) (*models.Tenant, error) {
	// get tenant for info
	targetTenant, err := getTenant(tenantName)
	if err != nil {
		return nil, err
	}

	consulService := &consul_api.AgentServiceRegistration{
		ID:      tenantName + "-" + service.Type,
		Name:    tenantName,
		Port:    service.Port,
		Address: service.Address,
	}

	doer := &commander{}

	// register the service to consul
	cmdConsulServiceReg := &consulServiceRegistrationCommand{
		consul:  consulClient,
		service: consulService,
	}
	doer.addCommand(cmdConsulServiceReg)

	// create grafana dashboard
	dashboardKey := dashboardRoot + service.Type
	dashboardInfo := &models.GrafanaDashboard{}
	cmdGrafanaDashboardCreation := &grafanaDashboardCreationCommand{
		consul:        consulClient,
		grafana:       grafanaOrgClients[targetTenant.GrafanaOrgID],
		tenant:        targetTenant,
		dashboardKey:  dashboardKey,
		dashboardInfo: dashboardInfo,
	}
	doer.addCommand(cmdGrafanaDashboardCreation)

	// link the service id to the tenant in the kv store
	newServiceEntry := models.Service{
		ServiceID:   consulService.ID,
		ServiceName: consulService.Name,
		Dashboard:   dashboardInfo,
	}
	cmdLinkService := &consulLinkServiceCommand{
		consul:     consulClient,
		tenantName: tenantName,
		newService: &newServiceEntry,
	}
	doer.addCommand(cmdLinkService)

	err = doer.execute()
	if err != nil {
		doer.rollback()
		return nil, err
	}

	return getTenant(tenantName)
}

// DeregisterService deregister a service from a tenant
//func DeregisterService(serviceID string, tenantName string) error {
//	// deregister the service from consul
//	log.Println("deregister service: " + serviceID)
//	err := consulClient.Agent().ServiceDeregister(serviceID)
//	if err != nil {
//		return err
//	}
//
//	// unlink the service id from the tenant in the kv store
//	log.Println("unlink service from " + tenantName)
//	path := tenantsRoot + tenantName
//
//	pair, _, err := consulClient.KV().Get(path, nil)
//	if err != nil {
//		return err
//	}
//
//	var tenantInStore almond.Tenant
//	err = json.Unmarshal(pair.Value, &tenantInStore)
//	if err != nil {
//		return err
//	}
//
//	serviceIDs := tenantInStore.ServiceIDs
//	var index int
//	for idx, id := range (serviceIDs) {
//		if id == serviceID {
//			index = idx
//			break
//		}
//	}
//	log.Println("service found at index " + string(index))
//
//	copy(serviceIDs[index:], serviceIDs[index+1:])
//	tenantInStore.ServiceIDs = serviceIDs[:len(tenantInStore.ServiceIDs)-1]
//
//	data, err := json.Marshal(tenantInStore)
//	if err != nil {
//		return err
//	}
//
//	p := &api.KVPair{Key: path, Value: data}
//	_, err = consulClient.KV().Put(p, nil)
//	return err
//}
