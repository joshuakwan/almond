package consul

import (
	"github.com/hashicorp/consul/api"
	"log"
	"github.com/astaxie/beego"
)

var (
	consulUrl    = beego.AppConfig.String(beego.AppConfig.String("runmode") + "::consul_url")
	consulClient = initializeConsulClient()
)

func initializeConsulClient() *api.Client {
	log.Println("Initializing consul API client")
	client, err := api.NewClient(&api.Config{Address: consulUrl})
	if err != nil {
		panic(err)
	}
	log.Println(client)
	return client
}
