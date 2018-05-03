package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:tenant/services`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"],
		beego.ControllerComments{
			Method: "RegisterGrafanaDashboard",
			Router: `/dashboards/:name`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"],
		beego.ControllerComments{
			Method: "CreateTenant",
			Router: `/tenants`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers:AlmondController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
