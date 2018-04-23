package consul

import (
	"github.com/astaxie/beego"
	"log"
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"github.com/joshuakwan/almond/models/common"
)

type ConsulController struct {
	beego.Controller
}

func (c *ConsulController) GetAll() {

}

func (c *ConsulController) Get() {

}

// @Title Post
// @Description register a service
// @router / [post]
func (c *ConsulController) Post() {
	var newService api.AgentServiceRegistration

	body := c.Ctx.Input.RequestBody
	log.Println(string(body))

	err := json.Unmarshal(body, &newService)

	if err !=nil {
		c.CustomAbort(400,"Bad request")
	}

	//reg := api.AgentServiceRegistration{
	//	ID:   newService.ID,
	//	Name: newService.Name,
	//	Port: newService.Port,
	//}
	err = consulClient.Agent().ServiceRegister(&newService)

	if err != nil {
		c.Data["json"] = common.Message{Text: "Service registration failed"}

	} else {
		c.Data["json"] = common.Message{Text: "Service registration succeeded"}
	}

	c.ServeJSON()
}

// @Title Delete
// @Description deregister a service
// @router /:id [delete]
func (c *ConsulController) Delete() {
	id := c.GetString(":id")

	err := consulClient.Agent().ServiceDeregister(id)
	if err != nil {
		c.Data["json"] = common.Message{Text: "Service deregistration failed"}

	} else {
		c.Data["json"] = common.Message{Text: "Service deregistration succeeded"}
	}

	c.ServeJSON()
}
