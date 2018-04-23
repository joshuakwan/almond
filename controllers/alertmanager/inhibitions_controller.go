package alertmanager

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/alertmanager"
	"log"
	"encoding/json"
	"strconv"
	"github.com/joshuakwan/almond/models/common"
	"github.com/prometheus/alertmanager/config"
)

type InhibitionController struct {
	beego.Controller
}

// @Title GetAll
// @Description get inhibition settings
// @router / [get]
func (i *InhibitionController) GetAll() {
	i.Data["json"] = liveConfig.InhibitRules
	i.ServeJSON()
}

// @Title Post
// @Description add an inhibition rule
// @router / [post]
func (i *InhibitionController) Post() {
	currentConfig := liveConfig.InhibitRules

	var newRule config.InhibitRule

	body := i.Ctx.Input.RequestBody
	log.Println(string(body))

	err :=json.Unmarshal(body, &newRule)
	if err!=nil{
		i.CustomAbort(400, "Invalid JSON object")
	}
	log.Println(&newRule)

	liveConfig.InhibitRules = alertmanager.AddInhibitRule(currentConfig, &newRule)
	log.Println(liveConfig.InhibitRules)

	go refreshAlertmanager()

	i.Data["json"] = liveConfig.InhibitRules
	i.ServeJSON()
}

// @Title Delete
// @Description delete an inhibition rule at a certain index (start at 0)
// @router /:index [delete]
func (i *InhibitionController) Delete() {
	currentConfig := liveConfig.InhibitRules

	index, err := strconv.Atoi(i.GetString(":index"))
	if err != nil {
		i.Data["json"] = common.Message{Text: "Invalid index value " + i.GetString(":index")}
	} else {
		rules, err := alertmanager.RemoveInhibitRule(currentConfig, index)
		if err != nil {
			i.Data["json"] = common.Message{Text: "Index " + string(index) + " not in the right range"}
		} else {
			liveConfig.InhibitRules = rules
			log.Println(liveConfig.InhibitRules)
			go refreshAlertmanager()
			i.Data["json"] = liveConfig.InhibitRules
		}
	}

	i.ServeJSON()
}
