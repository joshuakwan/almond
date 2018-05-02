package prometheus

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/facade"
)

var (
	configFilename = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::prometheus_config")
	config         = facade.GetPrometheusConfig(configFilename)
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
