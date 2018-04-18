package prometheus

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/models/prometheus"
)

var configFilename = beego.AppConfig.String("prometheus_config")

var config = getTotalConfig()

func getTotalConfig() *prometheus.Config {
	log.Println("Read prometheus configuration from " + configFilename)
	cfg, err := prometheus.LoadConfigFromFile(configFilename)
	if err != nil {
		panic(err)
	}
	return cfg
}
