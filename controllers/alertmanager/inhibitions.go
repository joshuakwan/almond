package alertmanager

import "github.com/astaxie/beego"

type InhibitionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get inhibition settings
// @router / [get]
func (i *InhibitionController) GetAll() {
	i.Data["json"] = config.InhibitRules
	i.ServeJSON()
}
