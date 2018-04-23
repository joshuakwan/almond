package facade

import (
	"net/http"
	prom_alertmanager "github.com/prometheus/alertmanager/config"
	"log"
)

func reloadPrometheusService(serviceUrl string) {
	url := serviceUrl + "/-/reload"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err!=nil {
		log.Fatalln(serviceUrl + " cannot be reloaded")
	} else {
		if resp.StatusCode == 200 {
			log.Println(serviceUrl + " reloaded")
		} else {
			log.Fatalln(serviceUrl + " fails to get reloaded")
		}
	}
}

func RefreshAlertmanager(url string, config *prom_alertmanager.Config,configFilename string) {
	writeAlertmanagerConfigToDisk(config, configFilename)
	reloadPrometheusService(url)
}