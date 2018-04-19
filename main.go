package main

import (
	_ "github.com/joshuakwan/almond/routers"

	"github.com/astaxie/beego"
	"log"
)

func main() {
	log.Println("Runmode: " + beego.AppConfig.String("runmode"))
	beego.Run()
}
