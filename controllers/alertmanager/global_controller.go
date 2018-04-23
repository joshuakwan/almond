package alertmanager

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/prometheus/alertmanager/config"
	"log"
	"github.com/joshuakwan/almond/models/alertmanager"
	"github.com/joshuakwan/almond/facade"
)

type GlobalController struct {
	beego.Controller
}

// @Title GetAll
// @Description get global settings
// @router / [get]
func (g *GlobalController) GetAll() {
	g.Data["json"] = liveConfig.Global
	g.ServeJSON()
}

// @Title Put
// @Description update global settings, partial update is supported
// @router / [put]
func (g *GlobalController) Put() {
	currentConfig := liveConfig.Global
	var newGlobal config.GlobalConfig

	body := g.Ctx.Input.RequestBody
	log.Println(string(body))

	err := json.Unmarshal(body, &newGlobal)
	if err!=nil{
		g.CustomAbort(400, "Invalid JSON object")
	}

	log.Println(newGlobal)

	alertmanager.Update(currentConfig, &newGlobal)
	log.Println(currentConfig)

	go facade.RefreshAlertmanager(alertmanagerUrl,liveConfig,configFilename)

	g.Data["json"] = currentConfig
	g.ServeJSON()
}

// @Title Delete
// @Description delete an item in global settings with the specified key
// @router /:key [delete]
func (g *GlobalController) Delete() {
	currentConfig := liveConfig.Global

	key := g.GetString(":key")
	log.Println(key)

	alertmanager.Delete(currentConfig, key)

	go facade.RefreshAlertmanager(alertmanagerUrl,liveConfig,configFilename)

	g.Data["json"] = currentConfig
	g.ServeJSON()
}
