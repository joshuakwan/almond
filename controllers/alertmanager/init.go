package alertmanager

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/alertmanager"
)

var configFilename = beego.AppConfig.String("alertmanager_config")

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
