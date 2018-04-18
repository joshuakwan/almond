package prometheus

import (
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

// @Title Get
// @Description Get current configuration
// @router / [get]
func (c *ConfigController) Get() {
	c.Data["json"] = config
	c.ServeJSON()
}
