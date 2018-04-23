package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/consul:ConsulController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/consul:ConsulController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/consul:ConsulController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/consul:ConsulController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
