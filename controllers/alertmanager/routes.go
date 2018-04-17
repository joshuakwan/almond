package alertmanager

import "github.com/astaxie/beego"

type RouteController struct {
	beego.Controller
}

// @Title GetAll
// @Description get route settings
// @router / [get]
func (r *RouteController) GetAll() {
	r.Data["json"] = config.Route
	r.ServeJSON()
}
