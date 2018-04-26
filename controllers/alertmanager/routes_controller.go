package alertmanager

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/alertmanager"
	"log"
	"encoding/json"
	"strconv"
	"github.com/joshuakwan/almond/models/common"
	"github.com/prometheus/alertmanager/config"
	"github.com/joshuakwan/almond/facade"
)

type RouteController struct {
	beego.Controller
}

// @Title GetAll
// @Description get route settings
// @router / [get]
func (r *RouteController) GetAll() {
	r.Data["json"] = liveConfig.Route
	r.ServeJSON()
}

// @Title Post
// @Description add a new sub route (now supports only layer 1)
// @router / [post]
func (r *RouteController) Post() {
	currentConfig := liveConfig.Route

	var newSubroute config.Route

	body:= r.Ctx.Input.RequestBody
	log.Println(string(body))

	err:=json.Unmarshal(body,&newSubroute)
	if err!=nil{
		r.CustomAbort(400, "Invalid JSON object")
	}
	log.Println(&newSubroute)

	liveConfig.Route = alertmanager.AddSubroute(currentConfig, &newSubroute)
	log.Println(liveConfig.Route)

	go facade.RefreshAlertmanager(alertmanagerUrl,liveConfig,configFilename)

	r.Data["json"] = liveConfig.Route
	r.ServeJSON()
}

// @Title Delete
// @Description delete a sub route at a certain index (starts at 0, now supports only layer 1)
// @router /:index [delete]
func (r *RouteController) Delete() {
	currentConfig := liveConfig.Route

	index, err := strconv.Atoi(r.GetString(":index"))
	if err != nil {
		r.Data["json"] = common.Message{Text: "Invalid index value " + r.GetString(":index")}
	} else {
		route, err := alertmanager.RemoveSubroute(currentConfig, index)
		if err != nil {
			r.Data["json"] = common.Message{Text: "Index " + string(index) + " not in the right range"}
		} else {
			liveConfig.Route = route
			log.Println(liveConfig.Route)
			go facade.RefreshAlertmanager(alertmanagerUrl,liveConfig,configFilename)
			r.Data["json"] = liveConfig.Route
		}
	}

	r.ServeJSON()
}
