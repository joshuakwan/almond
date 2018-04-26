// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/controllers/alertmanager"
	"github.com/joshuakwan/almond/controllers/prometheus"
	"github.com/joshuakwan/almond/controllers/almond"
)

func init() {
	namespaceRoot := "/api/v1"

	facadeNamespaceRoot := "/almond"
	nsFacade := beego.NewNamespace(namespaceRoot,
		beego.NSNamespace(facadeNamespaceRoot,
			beego.NSInclude(
				&almond.FacadeController{},
			),
		),
	)
	beego.AddNamespace(nsFacade)

	alertmanagerNamespaceRoot := "/alertmanager/config/"
	nsAlertmanager := beego.NewNamespace(namespaceRoot,
		beego.NSNamespace(alertmanagerNamespaceRoot,
			beego.NSInclude(
				&alertmanager.ConfigController{},
			),
		),
		beego.NSNamespace(alertmanagerNamespaceRoot+"global",
			beego.NSInclude(
				&alertmanager.GlobalController{},
			),
		),
		beego.NSNamespace(alertmanagerNamespaceRoot+"routes",
			beego.NSInclude(
				&alertmanager.RouteController{},
			),
		),
		beego.NSNamespace(alertmanagerNamespaceRoot+"inhibitions",
			beego.NSInclude(
				&alertmanager.InhibitionController{},
			),
		),
		beego.NSNamespace(alertmanagerNamespaceRoot+"receivers",
			beego.NSInclude(
				&alertmanager.ReceiverController{},
			),
		),
	)
	beego.AddNamespace(nsAlertmanager)

	prometheusNamespaceRoot := "/prometheus/config/"
	nsPrometheus := beego.NewNamespace(namespaceRoot,
		beego.NSNamespace(prometheusNamespaceRoot,
			beego.NSInclude(
				&prometheus.ConfigController{},
			),
		),
	)
	beego.AddNamespace(nsPrometheus)
}
