package alertmanager

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/alertmanager"
	"encoding/json"
	"log"
)

type GlobalController struct {
	beego.Controller
}

// @Title GetAll
// @Description get global settings
// @router / [get]
func (g *GlobalController) GetAll() {
	g.Data["json"] = config.Global
	g.ServeJSON()
}

// @Title Put
// @Description update global settings, partial update is supported
// @router / [put]
func (g *GlobalController) Put() {
	currentConfig := config.Global
	var newGlobal alertmanager.Global

	body := g.Ctx.Input.RequestBody
	log.Println(string(body))

	json.Unmarshal(body, &newGlobal)
	log.Println(newGlobal)

	currentConfig.Update(&newGlobal)
	log.Println(currentConfig)

	go refreshAlertmanager()

	g.Data["json"] = currentConfig
	g.ServeJSON()
}

// @Title Delete
// @Description delete an item in global settings with the specified key
// @router /:key [delete]
func (g *GlobalController) Delete() {
	currentConfig := config.Global

	key := g.GetString(":key")
	log.Println(key)

	currentConfig.Delete(key)

	go refreshAlertmanager()

	g.Data["json"] = currentConfig
	g.ServeJSON()
}
