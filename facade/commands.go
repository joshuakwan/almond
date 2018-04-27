package facade

import (
	"github.com/joshuakwan/almond/models/almond"
	grafana_models "github.com/joshuakwan/grafana-client/models"
	grafana_api "github.com/joshuakwan/grafana-client/api"
	consul_api "github.com/hashicorp/consul/api"
	"log"
	"errors"
	"fmt"
)

type command interface {
	do() error
	undo() error
}

type commander struct {
	commands      []command
	rollbackIndex int
}

func (e *commander) addCommand(c command) {
	e.commands = append(e.commands, c)
}

func (e *commander) execute() error {
	for i, command := range e.commands {
		err := command.do()
		if err != nil {
			e.rollbackIndex = i - 1
			return err
		}
	}
	return nil
}

func (e *commander) rollback() []error {
	if e.rollbackIndex < 0 {
		return nil
	}

	var errs []error
	for i := e.rollbackIndex; i >= 0; i-- {
		err := e.commands[i].undo()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

type grafanaOrgCreationCommand struct {
	grafana    *grafana_api.Client
	grafanaURL string
	tenant     *almond.Tenant
}

func (c *grafanaOrgCreationCommand) do() error {
	log.Println("DO create new grafana org", c.tenant.GrafanaOrgName)
	c.tenant.GrafanaURL = c.grafanaURL
	if c.tenant.GrafanaOrgName == "" {
		c.tenant.GrafanaOrgName = c.tenant.Name
	}
	message, err := c.grafana.CreateOrganization(&grafana_models.GrafanaOrganization{
		Name: c.tenant.GrafanaOrgName,
	})
	if err != nil {
		return err
	}
	c.tenant.GrafanaOrgID = message.OrgID
	return nil
}

func (c *grafanaOrgCreationCommand) undo() error {
	log.Println("UNDO create new grafana org ", c.tenant.GrafanaOrgName)
	message, err := c.grafana.DeleteOrganization(c.tenant.GrafanaOrgID)
	log.Println(message.Message)
	return err
}

type grafanaUserCreationCommand struct {
	grafana *grafana_api.Client
	tenant  *almond.Tenant
}

func (c *grafanaUserCreationCommand) do() error {
	log.Println("DO create the org user", c.tenant.GrafanaOrgUser)
	messageUser, err := c.grafana.CreateGlobalUser(&grafana_models.User{
		Name:     c.tenant.GrafanaOrgUser,
		Login:    c.tenant.GrafanaOrgUser,
		Password: c.tenant.GrafanaOrgUserPassword,
	})
	log.Println(messageUser)
	c.tenant.GrafanaOrgUserID = messageUser.ID
	return err
}

func (c *grafanaUserCreationCommand) undo() error {
	log.Println("UNDO create the org user", c.tenant.GrafanaOrgUser)
	_, err := c.grafana.DeleteGlobalUser(c.tenant.GrafanaOrgUserID)
	return err
}

type grafanaUserAssignmentCommand struct {
	grafana *grafana_api.Client
	tenant  *almond.Tenant
}

func (c *grafanaUserAssignmentCommand) do() error {
	log.Println("DO assign the user to the org", c.tenant.GrafanaOrgName)
	_, err := c.grafana.AddOrganizationUser(
		c.tenant.GrafanaOrgID,
		c.tenant.GrafanaOrgUser,
		"Admin")
	return err
}

func (c *grafanaUserAssignmentCommand) undo() error {
	// TODO implement grafana user un-assignment function
	return nil
}

type grafanaDatasourceCreationCommand struct {
	grafana *grafana_api.Client
	tenant  *almond.Tenant
}

func (c *grafanaDatasourceCreationCommand) do() error {
	log.Println("DO create datasource in the org")
	datasource := grafana_models.Datasource{
		Name:   "prometheus",
		Type:   "prometheus",
		URL:    prometheusUrl,
		Access: "proxy",
	}
	messageDatasource, err := c.grafana.AdminCreateDatasource(c.tenant.GrafanaOrgID, &datasource)
	log.Println(messageDatasource)
	return err
}

func (c *grafanaDatasourceCreationCommand) undo() error {
	// TODO implement grafana datasource deletion function
	return nil
}

type putTenantCommand struct {
	consul *consul_api.Client
	tenant *almond.Tenant
}

func (c *putTenantCommand) do() error {
	log.Println("DO put tenant into consul")
	return putTenant(c.consul, c.tenant)
}

// TODO howto
func (c *putTenantCommand) undo() error {
	log.Println("UNDO put tenant into consul")
	return nil
}

type consulServiceRegistrationCommand struct {
	consul  *consul_api.Client
	service *consul_api.AgentServiceRegistration
}

func (c *consulServiceRegistrationCommand) do() error {
	log.Println("DO register service:", c.service.ID)
	return c.consul.Agent().ServiceRegister(c.service)
}

func (c *consulServiceRegistrationCommand) undo() error {
	log.Println("UNDO register service:", c.service.ID)
	return c.consul.Agent().ServiceDeregister(c.service.ID)
}

type grafanaDashboardCreationCommand struct {
	consul        *consul_api.Client
	grafana       *grafana_api.Client
	tenant        *almond.Tenant
	dashboardKey  string
	dashboardInfo *almond.GrafanaDashboard
}

func (c *grafanaDashboardCreationCommand) do() error {
	log.Println("DO create dashboard", c.dashboardKey, "for", c.tenant.Name)
	log.Println("fetch dashboard:", c.dashboardKey)
	dashboardData, err := getConsulKVData(c.consul, c.dashboardKey)
	if err != nil {
		return err
	}
	if dashboardData == nil {
		return errors.New(fmt.Sprintf("dashboard not found at", c.dashboardKey))
	}

	message, err := c.grafana.AdminCreateDashboardFromJSON(c.tenant.GrafanaOrgID, dashboardData)
	if err != nil {
		return err
	}
	c.dashboardInfo.UID = message.UID
	c.dashboardInfo.URL = message.URL
	c.dashboardInfo.Slug = message.Slug

	return nil
}

func (c *grafanaDashboardCreationCommand) undo() error {
	log.Println("UNDO create dashboard", c.dashboardKey, "for", c.tenant.Name)
	message, err := c.grafana.AdminDeleteDashboardByUID(c.tenant.GrafanaOrgID, c.dashboardInfo.UID)
	log.Println(message.Title)
	return err
}

type consulLinkServiceCommand struct {
	consul     *consul_api.Client
	tenantName string
	newService *almond.Service
}

func (c *consulLinkServiceCommand) do() error {
	log.Println("DO link service to", c.tenantName)
	tenantInStore, err := getTenant(c.tenantName)
	if err != nil {
		return err
	}

	tenantInStore.Services = append(tenantInStore.Services, c.newService)
	return putTenant(c.consul, tenantInStore)
}

// TODO
func (c *consulLinkServiceCommand) undo() error {
	return nil
}
