package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"],
		beego.ControllerComments{
			Method: "DeregisterService",
			Router: `/:tenant/:service`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"],
		beego.ControllerComments{
			Method: "CreateTenant",
			Router: `/tenants`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
