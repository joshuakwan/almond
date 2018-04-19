package alertmanager

import (
	"log"

	"github.com/joshuakwan/almond/models/alertmanager"
	"github.com/astaxie/beego"
	"net/http"
)

var (
	configFilename = beego.AppConfig.String(beego.AppConfig.String("runmode")+"::alertmanager_config")
	config = getTotalConfig()

	alertmanagerUrl = beego.AppConfig.String(beego.AppConfig.String("runmode")+"::alertmanager_url")
)

func getTotalConfig() *alertmanager.Config {
	log.Println("Read alertmanager configuration from " + configFilename)
	cfg, err := alertmanager.LoadConfigFromFile(configFilename)
	if err != nil {
		panic(err)
	}
	return cfg
}

func refreshAlertmanager() {
	writeTotalConfig()
	reloadAlertmanager()
}

func writeTotalConfig() {
	log.Println("Write alertmanager configuration to " + configFilename)
	err := alertmanager.SaveConfigToFile(config, configFilename)
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