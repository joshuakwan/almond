package alertmanager

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/alertmanager"
	"log"
	"encoding/json"
	"github.com/joshuakwan/almond/models/common"
	"github.com/prometheus/alertmanager/config"
	"github.com/joshuakwan/almond/facade"
)

type ReceiverController struct {
	beego.Controller
}

// @Title GetAll
// @Description get receiver settings
// @router / [get]
func (r *ReceiverController) GetAll() {
	r.Data["json"] = liveConfig.Receivers
	r.ServeJSON()
}

// @Title Post
// @Description add a new receiver
// @router / [post]
func (r *ReceiverController) Post() {
	currentConfig := liveConfig.Receivers
	log.Println(len(currentConfig))

	var newReceiver config.Receiver

	body := r.Ctx.Input.RequestBody
	log.Println(string(body))

	err := json.Unmarshal(body, &newReceiver)
	if err != nil {
		r.CustomAbort(400, "Invalid JSON object")
	}
	log.Println(&newReceiver)

	receivers, err := alertmanager.AddReceiver(currentConfig, &newReceiver)

	if err != nil {
		message := common.Message{Text: "Receiver " + newReceiver.Name + " already exists"}
		r.Data["json"] = message
	} else {
		liveConfig.Receivers = receivers
		go facade.RefreshAlertmanager(alertmanagerUrl, liveConfig, configFilename)
		r.Data["json"] = liveConfig.Receivers
	}

	r.ServeJSON()
}

// @Title Delete
// @Description delete a receiver by name
// @router /:name [delete]
func (r *ReceiverController) Delete() {
	currentConfig := liveConfig.Receivers

	name := r.GetString(":name")
	log.Println("receiver's name to delete: " + name)

	receivers, err := alertmanager.RemoveReceiver(currentConfig, name)

	if err != nil {
		message := common.Message{Text: "Receiver " + name + " not found"}
		r.Data["json"] = message
	} else {
		liveConfig.Receivers = receivers
		go facade.RefreshAlertmanager(alertmanagerUrl, liveConfig, configFilename)
		r.Data["json"] = liveConfig.Receivers
	}

	r.ServeJSON()
}
