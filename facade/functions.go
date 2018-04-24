package facade

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/almond"
	"github.com/hashicorp/consul/api"
	grafana_api "github.com/joshuakwan/grafana-client/api"
	grafana_models "github.com/joshuakwan/grafana-client/models"
	"encoding/json"
	"log"
	"errors"
)

var (
	consulUrl    = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::consul_url")
	ConsulClient = InitializeConsulClient(consulUrl)

	grafanaUrl      = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_url")
	grafanaKey      = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_bearer_token")
	grafanaUser     = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_admin_user")
	grafanaPassword = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_admin_password")
	grafanaClient   = grafana_api.Client{GrafanaURL: grafanaUrl,
		BearerToken: grafanaKey,
		AdminUser: grafanaUser,
		AdminPassword: grafanaPassword}

	tenantsRoot   = "tenants/"
	dashboardRoot = "grafana_dashboards/"
	//servicesKey = "/services"
)

// RegisterDashboard register a dashboard (its json data) to consul under grafana_dashboards/
func RegisterDashboard(name string, jsondata []byte) error {
	key := dashboardRoot + name
	log.Println("register dashboard at " + key)

	p := &api.KVPair{Key: key, Value: jsondata}
	_, err := ConsulClient.KV().Put(p, nil)
	return err
}

// 3 things to do:
//   1. check if the tenant already exists
//   2. create a new grafana org
//   3. create a new folder in consul kv store tenants/{tenant name}
func CreateTenant(tenant *almond.Tenant) (*almond.Tenant, error) {
	key := tenantsRoot + tenant.Name

	// check if the tenant already exists
	log.Println("check if tenant " + tenant.Name + " exists")
	pair, _, err := ConsulClient.KV().Get(key, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if pair != nil {
		log.Println("tenant " + tenant.Name + " already exists")
		return nil, errors.New("tenant " + tenant.Name + " already exists")
	}

	// create a new grafana org
	log.Println("create tenant " + tenant.Name)
	tenant.GrafanaURL = grafanaUrl
	if tenant.GrafanaOrgName == "" {
		tenant.GrafanaOrgName = tenant.Name
	}
	log.Println("create new grafana org: " + tenant.GrafanaOrgName)
	newOrg := grafana_models.GrafanaOrganization{Name: tenant.GrafanaOrgName}
	_, err = grafanaClient.CreateOrganization(&newOrg)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// create a new tenant entry in consul kv store as tenants/{tenant name}
	value, err := json.Marshal(tenant)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	p := &api.KVPair{Key: key, Value: value}
	_, err = ConsulClient.KV().Put(p, nil)
	return tenant, err
}

// RegisterService register a service to a tenant
// 3 parts:
//   1. register the service to consul
//   2. do grafana stuff [TODO] - use service.Name to indicate the dashboard type
//   3. link the service id to the tenant in the kv store
func RegisterService(tenantName string, service *api.AgentServiceRegistration) (*almond.Tenant, error) {
	// register the service to consul
	log.Println("register service: " + service.ID)
	err := ConsulClient.Agent().ServiceRegister(service)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// create grafana dashboard
	dashboardKey := dashboardRoot + service.Name

	// fetch dashboard
	log.Println("fetch dashboard: " + dashboardKey)
	pair, _, err := ConsulClient.KV().Get(dashboardKey, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if pair == nil {
		return nil, errors.New("dashboard not found at " + dashboardKey)
	}

	message, err := grafanaClient.CreateDashboardFromJSON(pair.Value)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	dashboardInfo := almond.GrafanaDashboard{UID: message.UID,
		URL: message.URL,
		Slug: message.Slug}

	// link the service id to the tenant in the kv store
	log.Println("link service to " + tenantName)
	path := tenantsRoot + tenantName

	pair, _, err = ConsulClient.KV().Get(path, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var tenantInStore almond.Tenant
	err = json.Unmarshal(pair.Value, &tenantInStore)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	newServiceEntry := almond.Service{ServiceID: service.ID,
		ServiceName: service.Name,
		Dashboard: &dashboardInfo}

	tenantInStore.Services = append(tenantInStore.Services, &newServiceEntry)
	data, err := json.Marshal(tenantInStore)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	p := &api.KVPair{Key: path, Value: data}
	_, err = ConsulClient.KV().Put(p, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &tenantInStore, err
}

// DeregisterService deregister a service from a tenant
//func DeregisterService(serviceID string, tenantName string) error {
//	// deregister the service from consul
//	log.Println("deregister service: " + serviceID)
//	err := ConsulClient.Agent().ServiceDeregister(serviceID)
//	if err != nil {
//		return err
//	}
//
//	// unlink the service id from the tenant in the kv store
//	log.Println("unlink service from " + tenantName)
//	path := tenantsRoot + tenantName
//
//	pair, _, err := ConsulClient.KV().Get(path, nil)
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
//	_, err = ConsulClient.KV().Put(p, nil)
//	return err
//}
