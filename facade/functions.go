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
	"net/http"
)

var (
	consulUrl    = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::consul_url")
	ConsulClient = InitializeConsulClient(consulUrl)

	prometheusUrl = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::prometheus_url")

	grafanaUrl      = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_url")
	grafanaKey      = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_bearer_token")
	grafanaUser     = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_admin_user")
	grafanaPassword = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_admin_password")
	grafanaClient   = grafana_api.Client{
		GrafanaURL:    grafanaUrl,
		BearerToken:   grafanaKey,
		AdminUser:     grafanaUser,
		AdminPassword: grafanaPassword,
	}

	consulRoot    = "almond/"
	tenantsRoot   = consulRoot + "tenants/"
	dashboardRoot = consulRoot + "grafana_dashboards/"
	//servicesKey = "/services"

	grafanaDatasource *grafana_models.Datasource
)

func init() {
	log.Println("do some ugly initialization stuff")
	log.Println("check the liveness of dependent services")

	checkGrafana()
	checkConsul()
}

func checkGrafana() {
	log.Println("check the readiness of grafana, do some stuff as needed")
	// grafana should be living
	err := checkURLLiveness(grafanaUrl)
	if err != nil {
		log.Println("grafana unreachable, error: ", err)
		panic(err)
	} else {
		log.Println("grafana is running")
	}
}

func checkConsul() {
	log.Println("check the readiness of consul, do some stuff as needed")
	err := checkURLLiveness(consulUrl)
	if err != nil {
		log.Println("consul unreachable, error: ", err)
		panic(err)
	} else {
		log.Println("consul is running")
	}
}

func checkURLLiveness(url string) error {
	_, err := http.Get(url)
	return err
}

// RegisterDashboard register a dashboard (its json data) to consul under grafana_dashboards/
func RegisterDashboard(name string, jsondata []byte) error {
	key := dashboardRoot + name
	log.Println("register dashboard at " + key)

	p := &api.KVPair{Key: key, Value: jsondata}
	_, err := ConsulClient.KV().Put(p, nil)
	return err
}

func getData(key string) ([]byte, error) {
	pair, _, err := ConsulClient.KV().Get(key, nil)
	if pair == nil {
		return nil, errors.New(key + " not found")
	}
	return pair.Value, err
}

func getTenant(tenantName string) (*almond.Tenant, error) {
	data, err := getData(tenantsRoot + tenantName)
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

func checkIfTenantExists(tenantName string) bool {
	tenant, _ := getTenant(tenantName)

	if tenant == nil {
		return false
	} else {
		return true
	}
}

func putTenant(tenant *almond.Tenant) error {
	data, err := json.Marshal(tenant)
	if err != nil {
		return err
	}

	p := &api.KVPair{Key: tenantsRoot + tenant.Name, Value: data}
	_, err = ConsulClient.KV().Put(p, nil)
	return err
}

func getDashboardData(dashboardKey string) ([]byte, error) {
	return getData(dashboardKey)
}

// things to do:
//   1. check if the tenant already exists
//   2. create a new grafana org
//   3. create a new org user
//   4. assign the user to the new org
//   5. TODO generate an admin key of the new org - need to redesign, defer to post PoC phase
//   6. TODO create datasource
//   7. create a new folder in consul kv store tenants/{tenant name}
// TODO need a fallback mechanism - defer to post PoC phase
// TODO happy path done, need to make it robust and consistent
func CreateTenant(tenant *almond.Tenant) (*almond.Tenant, error) {
	// 1. check if the tenant already exists
	log.Println("check if tenant " + tenant.Name + " exists")
	if checkIfTenantExists(tenant.Name) == true {
		return nil, errors.New("tenant " + tenant.Name + " already exists")
	}

	// 2. create a new grafana org
	log.Println("create tenant " + tenant.Name)
	tenant.GrafanaURL = grafanaUrl
	if tenant.GrafanaOrgName == "" {
		tenant.GrafanaOrgName = tenant.Name
	}
	log.Println("create new grafana org: " + tenant.GrafanaOrgName)
	message, err := grafanaClient.CreateOrganization(&grafana_models.GrafanaOrganization{
		Name: tenant.GrafanaOrgName,
	})
	if err != nil {
		return nil, err
	}

	tenant.GrafanaOrgID = message.OrgID

	// 3. create a new org user
	log.Println("create the org user " + tenant.GrafanaOrgName)
	messageUser, err := grafanaClient.CreateGlobalUser(&grafana_models.User{
		Name:     tenant.GrafanaOrgUser,
		Login:    tenant.GrafanaOrgUser,
		Password: tenant.GrafanaOrgUserPassword,
	})
	log.Println(messageUser)
	if err != nil {
		return nil, err
	}

	// 4. assign the user to the new org
	log.Println("assign the user to the org " + tenant.GrafanaOrgName)
	_, err = grafanaClient.AddOrganizationUser(tenant.GrafanaOrgID, tenant.GrafanaOrgUser, "Admin")
	if err != nil {
		return nil, err
	}

	// 6. create datasource
	log.Println("create datasource in the org")
	datasource := grafana_models.Datasource{
		Name:   tenant.Name + "_prometheus",
		Type:   "prometheus",
		URL:    prometheusUrl,
		Access: "proxy",
	}
	messageDatasource, err := grafanaClient.AdminCreateDatasource(tenant.GrafanaOrgID, &datasource)
	log.Println(messageDatasource)
	if err != nil {
		return nil, err
	}

	// 7. create a new tenant entry in consul kv store as tenants/{tenant name}
	err = putTenant(tenant)
	if err != nil {
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
func RegisterService(tenantName string, service *api.AgentServiceRegistration) (*almond.Tenant, error) {
	// get tenant for info
	targetTenant, err := getTenant(tenantName)
	if err != nil {
		return nil, err
	}

	// register the service to consul
	log.Println("register service: " + service.ID)
	err = ConsulClient.Agent().ServiceRegister(service)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// create grafana dashboard
	dashboardKey := dashboardRoot + service.Name

	// fetch dashboard
	log.Println("fetch dashboard: " + dashboardKey)
	dashboardData, err := getDashboardData(dashboardKey)
	if err != nil {
		return nil, err
	}
	if dashboardData == nil {
		return nil, errors.New("dashboard not found at " + dashboardKey)
	}

	message, err := grafanaClient.AdminCreateDashboardFromJSON(targetTenant.GrafanaOrgID, dashboardData)
	if err != nil {
		return nil, err
	}
	dashboardInfo := almond.GrafanaDashboard{
		UID:  message.UID,
		URL:  message.URL,
		Slug: message.Slug,
	}

	// link the service id to the tenant in the kv store
	log.Println("link service to " + tenantName)

	tenantInStore, err := getTenant(tenantName)
	if err != nil {
		return nil, err
	}

	newServiceEntry := almond.Service{
		ServiceID:   service.ID,
		ServiceName: service.Name,
		Dashboard:   &dashboardInfo,
	}

	tenantInStore.Services = append(tenantInStore.Services, &newServiceEntry)

	err = putTenant(tenantInStore)
	if err != nil {
		return nil, err
	}

	return tenantInStore, nil
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
