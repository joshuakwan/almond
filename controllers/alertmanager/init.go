package alertmanager

import (
	"log"

	"github.com/joshuakwan/almond/models/alertmanager"
	"github.com/prometheus/alertmanager/config"
	"github.com/astaxie/beego"
	"net/http"
)

var (
	configFilename = beego.AppConfig.String(beego.AppConfig.String("runmode")+"::alertmanager_config")
	liveConfig = getTotalConfig()

	alertmanagerUrl = beego.AppConfig.String(beego.AppConfig.String("runmode")+"::alertmanager_url")
)

func getTotalConfig() *config.Config {
	log.Println("Read alertmanager configuration from " + configFilename)
	cfgLoad, _, err := config.LoadFile(configFilename)
	if err != nil {
		panic(err)
	}
	return cfgLoad
}

func refreshAlertmanager() {
	writeTotalConfig()
	reloadAlertmanager()
}

func writeTotalConfig() {
	log.Println("Write alertmanager configuration to " + configFilename)
	err := alertmanager.SaveConfigToFile(liveConfig, configFilename)
	if err != nil {
		panic(err)
	}
}

func reloadAlertmanager() {
	url := alertmanagerUrl + "/-/reload"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err!=nil {
		log.Fatalln(alertmanagerUrl + " cannot be reloaded")
	} else {
		if resp.StatusCode == 200 {
			log.Println(alertmanagerUrl + " reloaded")
		} else {
			log.Fatalln(alertmanagerUrl + " fails to get reloaded")
		}
	}
}