package facade

import (
	"net/http"
	prom_alertmanager "github.com/prometheus/alertmanager/config"
	"log"
	"github.com/joshuakwan/almond/utils"
)

func checkGrafana() {
	log.Println("check the readiness of grafana, do some stuff as needed")
	// grafana should be living
	err := utils.CheckURLLiveness(grafanaUrl)
	if err != nil {
		log.Println("grafana unreachable, error: ", err)
		panic(err)
	} else {
		log.Println("grafana is running at", grafanaUrl)
	}
}

func checkConsul() {
	log.Println("check the readiness of consul, do some stuff as needed")
	err := utils.CheckURLLiveness(consulUrl)
	if err != nil {
		log.Println("consul unreachable, error:", err)
		panic(err)
	} else {
		log.Println("consul is running at", consulUrl)
	}
}

func checkPrometheus() {
	log.Println("check the readiness of prometheus, do some stuff as needed")
	err := utils.CheckURLLiveness(prometheusUrl)
	if err != nil {
		log.Println("prometheus unreachable, error:", err)
		panic(err)
	} else {
		log.Println("prometheus is running at", prometheusUrl)
	}
}

func checkAlertmanager() {
	log.Println("check the readiness of alertmanager, do some stuff as needed")
	err := utils.CheckURLLiveness(alertmanagerUrl)
	if err != nil {
		log.Println("alertmanager unreachable, error:", err)
		panic(err)
	} else {
		log.Println("alertmanager is running at", alertmanagerUrl)
	}
}

func reloadPrometheusService(serviceUrl string) {
	url := serviceUrl + "/-/reload"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Fatalln(serviceUrl + " cannot be reloaded")
	} else {
		if resp.StatusCode == 200 {
			log.Println(serviceUrl, "reloaded")
		} else {
			log.Fatalln(serviceUrl, "fails to get reloaded")
		}
	}
}

func RefreshAlertmanager(url string, config *prom_alertmanager.Config, configFilename string) {
	writeAlertmanagerConfigToDisk(config, configFilename)
	reloadPrometheusService(url)
}
