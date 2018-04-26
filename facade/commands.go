package facade

import (
	"github.com/joshuakwan/almond/models/almond"
	grafana_models "github.com/joshuakwan/grafana-client/models"
	grafana_api "github.com/joshuakwan/grafana-client/api"
	consul_api "github.com/hashicorp/consul/api"
	"log"
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
		Name:   c.tenant.Name + "_prometheus",
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