package controllers

import (
	"log"
	"github.com/joshuakwan/almond/utils"
	"github.com/astaxie/beego"
	grafana_models "github.com/joshuakwan/grafana-client/models"
)

// constants
const (
	consulRoot    = "almond/"
	tenantsRoot   = consulRoot + "tenants/"
	dashboardRoot = consulRoot + "grafana_dashboards/"
)

// HTTP constants
const (
	HTTP_CODE_OK          = 200
	HTTP_CODE_BAD_REQUEST = 400
	HTTP_CODE_ERROR       = 500
)

// URLs
var (
	consulUrl       = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::consul_url")
	grafanaUrl      = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::grafana_url")
	prometheusUrl   = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::prometheus_url")
	promOperatorUrl = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::prom_operator_url")
)

// clients
var (
	consulClient  = getConsulClient(consulUrl)
	grafanaClient = getGrafanaClient(grafanaUrl,
		beego.AppConfig.String(beego.AppConfig.String("runmode")+"::grafana_bearer_token"),
		beego.AppConfig.String(beego.AppConfig.String("runmode")+"::grafana_admin_user"),
		beego.AppConfig.String(beego.AppConfig.String("runmode")+"::grafana_admin_password"))

	// key = orgID
	grafanaOrgClients = getGrafanaOrganizationClients()
)

// variables
var (
	grafanaDatasource *grafana_models.Datasource
)

func checkPromOperator() {
	log.Println("check the readiness of prometheus operator")
	if err := utils.CheckURLLiveness(promOperatorUrl); err != nil {
		log.Println("prometheus operator unreachable, error: ", err)
		panic(err)
	}
	log.Println("prometheus operator is running at", promOperatorUrl)
}

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
	err := utils.CheckURLLiveness("http://" + consulUrl)
	if err != nil {
		log.Println("consul unreachable, error:", err)
		panic(err)
	} else {
		log.Println("consul is running at", consulUrl)
	}
}

func init() {
	log.Println("do some ugly initialization stuff")
	log.Println("check the liveness of dependent services")

	checkPromOperator()
	checkConsul()
	checkGrafana()
}
