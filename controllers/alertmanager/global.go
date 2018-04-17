package alertmanager

import "github.com/astaxie/beego"

type GlobalController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all settings
// @router / [get]
func (g *GlobalController) GetAll() {
	g.Data["json"] = "global"
	g.ServeJSON()
}
