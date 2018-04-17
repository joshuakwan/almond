package alertmanager

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/alertmanager"
)

type ConfigController struct {
	beego.Controller
}

// @Title Get
// @Description Get current configuration
// @router / [get]
func (c *ConfigController) Get() {
	filename := beego.AppConfig.String("alertmanager_config")
	config, err := alertmanager.LoadConfigFromFile(filename)
	if err != nil {
		c.Data["json"] = "error"
	}
	c.Data["json"] = config
	c.ServeJSON()
}
