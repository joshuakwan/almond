package facade

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/almond"
	"github.com/hashicorp/consul/api"
	"encoding/json"
	"log"
)

var (
	consulUrl    = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::consul_url")
	ConsulClient = InitializeConsulClient(consulUrl)

	tenantsRoot = "tenants/"
	servicesKey = "/services"
)

func CreateTenant(tenant *almond.Tenant) error {
	log.Println("create tenant " + tenant.Name)

	key := tenantsRoot + tenant.Name
	value, err := json.Marshal(tenant)
	if err != nil {
		return err
	}

	p := &api.KVPair{Key: key, Value: value}
	_, err = ConsulClient.KV().Put(p, nil)
	return err
}

// RegisterService register a service to a tenant
// 3 parts:
//   1. register the service to consul
//   2. link the service id to the tenant in the kv store
//   3. do grafana stuff [TODO]
// TODO consider extracting the steps to functions for parallel handling
func RegisterService(service *api.AgentServiceRegistration, tenant *almond.Tenant) error {
	// register the service to consul
	log.Println("register service: " + service.ID)
	err := ConsulClient.Agent().ServiceRegister(service)
	if err != nil {
		return err
	}

	// link the service id to the tenant in the kv store
	log.Println("link service to " + tenant.Name)
	path := tenantsRoot + tenant.Name

	pair, _, err := ConsulClient.KV().Get(path, nil)
	if err != nil {
		return err
	}

	var tenantInStore almond.Tenant
	err = json.Unmarshal(pair.Value, &tenantInStore)
	if err != nil {
		return err
	}

	tenantInStore.ServiceIDs = append(tenantInStore.ServiceIDs, service.ID)
	data, err := json.Marshal(tenantInStore)
	if err != nil {
		return err
	}

	p := &api.KVPair{Key: path, Value: data}
	_, err = ConsulClient.KV().Put(p, nil)
	return err
}

// DeregisterService deregister a service from a tenant
func DeregisterService(serviceID string, tenantName string) error {
	// deregister the service from consul
	log.Println("deregister service: " + serviceID)
	err := ConsulClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		return err
	}

	// unlink the service id from the tenant in the kv store
	log.Println("unlink service from " + tenantName)
	path := tenantsRoot + tenantName

	pair, _, err := ConsulClient.KV().Get(path, nil)
	if err != nil {
		return err
	}

	var tenantInStore almond.Tenant
	err = json.Unmarshal(pair.Value, &tenantInStore)
	if err != nil {
		return err
	}

	serviceIDs := tenantInStore.ServiceIDs
	var index int
	for idx, id := range (serviceIDs) {
		if id == serviceID {
			index = idx
			break
		}
	}
	log.Println("service found at index " + string(index))

	copy(serviceIDs[index:], serviceIDs[index+1:])
	tenantInStore.ServiceIDs = serviceIDs[:len(tenantInStore.ServiceIDs)-1]

	data, err := json.Marshal(tenantInStore)
	if err != nil {
		return err
	}

	p := &api.KVPair{Key: path, Value: data}
	_, err = ConsulClient.KV().Put(p, nil)
	return err
}
