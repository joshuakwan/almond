package alertmanager

import (
	"log"

	"github.com/joshuakwan/almond/models/alertmanager"
	"github.com/astaxie/beego"
)

var configFilename = beego.AppConfig.String(beego.AppConfig.String("runmode")+"::alertmanager_config")

var config = getTotalConfig()

func getTotalConfig() *alertmanager.Config {
	log.Println("Read alertmanager configuration from " + configFilename)
	cfg, err := alertmanager.LoadConfigFromFile(configFilename)
	if err != nil {
		panic(err)
	}
	return cfg
}

func writeTotalConfig() {
	log.Println("Write alertmanager configuration to " + configFilename)
	err := alertmanager.SaveConfigToFile(config, configFilename)
	if err != nil {
		panic(err)
	}
}
