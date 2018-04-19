package alertmanager

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/alertmanager"
	"log"
	"encoding/json"
	"strconv"
	"github.com/joshuakwan/almond/models/common"
)

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

// @Title Post
// @Description add an inhibition rule
// @router / [post]
func (i *InhibitionController) Post() {
	currentConfig := config.InhibitRules

	var newRule alertmanager.InhibitRule

	body := i.Ctx.Input.RequestBody
	log.Println(string(body))

	json.Unmarshal(body, &newRule)
	log.Println(&newRule)

	config.InhibitRules = alertmanager.AddInhibitRule(currentConfig, &newRule)
	log.Println(config.InhibitRules)

	go refreshAlertmanager()

	i.Data["json"] = config.InhibitRules
	i.ServeJSON()
}

// @Title Delete
// @Description delete an inhibition rule at a certain index (start at 0)
// @router /:index [delete]
func (i *InhibitionController) Delete() {
	currentConfig := config.InhibitRules

	index, err := strconv.Atoi(i.GetString(":index"))
	if err != nil {
		i.Data["json"] = common.Message{Text: "Invalid index value " + i.GetString(":index")}
	} else {
		rules, err := alertmanager.RemoveInhibitRule(currentConfig, index)
		if err != nil {
			i.Data["json"] = common.Message{Text: "Index " + string(index) + " not in the right range"}
		} else {
			config.InhibitRules = rules
			log.Println(config.InhibitRules)
			go refreshAlertmanager()
			i.Data["json"] = config.InhibitRules
		}
	}

	i.ServeJSON()
}
