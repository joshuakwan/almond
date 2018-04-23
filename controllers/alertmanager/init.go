package alertmanager

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/facade"
)

var (
	alertmanagerUrl = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::alertmanager_url")
	configFilename  = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::alertmanager_config")
	liveConfig      = facade.GetAlertmanagerConfig(configFilename)
)
