package facade

import (
	"log"
	prom_alertmanager "github.com/prometheus/alertmanager/config"
	prom_prometheus "github.com/joshuakwan/almond/models/prometheus"
	"github.com/joshuakwan/almond/models/alertmanager"
)

func GetAlertmanagerConfig(configFilename string) *prom_alertmanager.Config {
	log.Println("Read alertmanager configuration from " + configFilename)
	cfg, _, err := prom_alertmanager.LoadFile(configFilename)
	if err != nil {
		panic(err)
	}
	return cfg
}

func writeAlertmanagerConfigToDisk(config *prom_alertmanager.Config, configFilename string) {
	log.Println("Write alertmanager configuration to " + configFilename)
	err := alertmanager.SaveConfigToFile(config, configFilename)
	if err != nil {
		panic(err)
	}
}

func GetPrometheusConfig(configFilename string) *prom_prometheus.Config {
	log.Println("Read prometheus configuration from " + configFilename)
	cfg, err := prom_prometheus.LoadFile(configFilename)
	if err != nil {
		panic(err)
	}
	return cfg
}

func writePrometheusConfigToDisk(config *prom_prometheus.Config, configFilename string) {
	log.Println("Write prometheus configuration to " + configFilename)
	err := prom_prometheus.SaveConfigToFile(config, configFilename)
	if err != nil {
		panic(err)
	}
}
