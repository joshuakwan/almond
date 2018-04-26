package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:tenant/services`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/almond:FacadeController"],
		beego.ControllerComments{
			Method: "RegisterGrafanaDashboard",
			Router: `/dashboards/:name`,
			AllowHTTPMethods: []string{"post"},
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
