package facade

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func InitializeConsulClient(url string) *api.Client {
	log.Println("Initializing consul API client")
	client, err := api.NewClient(&api.Config{Address: url})
	if err != nil {
		panic(err)
	}
	log.Println(client)
	return client
}
