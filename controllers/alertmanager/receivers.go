package alertmanager

import "github.com/astaxie/beego"

type ReceiverController struct {
	beego.Controller
}

// @Title GetAll
// @Description get receiver settings
// @router / [get]
func (r *ReceiverController) GetAll() {
	r.Data["json"] = config.Receivers
	r.ServeJSON()
}
