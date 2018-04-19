package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ConfigController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ConfigController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:GlobalController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:GlobalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:GlobalController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:GlobalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:GlobalController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:GlobalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:key`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:InhibitionController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:InhibitionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:InhibitionController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:InhibitionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:InhibitionController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:InhibitionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:index`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ReceiverController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ReceiverController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ReceiverController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ReceiverController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ReceiverController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:ReceiverController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:name`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:RouteController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:RouteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:RouteController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:RouteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:RouteController"] = append(beego.GlobalControllerRouter["github.com/joshuakwan/almond/controllers/alertmanager:RouteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:index`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
