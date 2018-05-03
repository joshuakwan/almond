package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"encoding/json"
	"github.com/joshuakwan/almond/models"
)

type AlmondController struct {
	beego.Controller
}

// @Title Test
// @Description for development manual test purpose
// @router /test [post]
func (c *AlmondController) Test() {

}

// @Title CreateTenant
// @Description create a new tenant
// Sample request:
// {
//	    "name":"org",
//		"grafana_org": "org-org",
//		"grafana_org_user": "org-admin",
//		"grafana_org_user_password": "Passw0rd",
//		"description":"yet another noisy tenant"
//}
// @router /tenants [post]
func (c *AlmondController) CreateTenant() {
	body := c.Ctx.Input.RequestBody
	log.Println(string(body))

	var newTenant models.Tenant
	err := json.Unmarshal(body, &newTenant)
	if err != nil {
		c.CustomAbort(HTTP_CODE_BAD_REQUEST, err.Error())
	}

	tenant, err := CreateTenant(&newTenant)
	if err != nil {
		c.CustomAbort(HTTP_CODE_BAD_REQUEST, err.Error())
	}

	c.Data["json"] = tenant
	c.ServeJSON()
}

// @Title Post
// @Description add a new service of a tenant to monitoring and alerting
// Sample request:
// {
//    	"id": "tenant-redis",
//    	"name": "redis",
//    	"port": 9121,
//    	"address": "127.0.0.1"
// }
// @router /:tenant/services [post]
func (c *AlmondController) Post() {
	tenantName := c.GetString(":tenant")
	body := c.Ctx.Input.RequestBody
	log.Println("register service to", tenantName)
	log.Println(string(body))

	var newServiceReg models.ServiceRegistration
	err := json.Unmarshal(body, &newServiceReg)
	if err != nil {
		c.CustomAbort(HTTP_CODE_BAD_REQUEST, err.Error())
	}

	tenant, err := RegisterService(tenantName, &newServiceReg)
	if err != nil {
		c.CustomAbort(HTTP_CODE_BAD_REQUEST, err.Error())
	}

	c.Data["json"] = tenant
	c.ServeJSON()
}

// @Title DeregisterService
// @Description deregister a service from a tenant
// @router /:tenant/:service [delete]
//func (f *AlmondController) DeregisterService() {
//	tenant := f.GetString(":tenant")
//	service := f.GetString(":service")
//	log.Println("deregister " + service + " from " + tenant)
//	err := facade.DeregisterService(service,tenant)
//	if err != nil {
//		f.CustomAbort(400, "Fail to deregister service from the tenant")
//	}
//}

// @Title RegisterGrafanaDashboard
// @Description register a grafana dashboard
// @router /dashboards/:name [post]
func (c *AlmondController) RegisterGrafanaDashboard() {
	body := c.Ctx.Input.RequestBody
	name := c.GetString(":name")
	log.Println(string(body))

	err := RegisterDashboard(name, body)

	if err != nil {
		c.CustomAbort(HTTP_CODE_BAD_REQUEST, err.Error())
	}
}
