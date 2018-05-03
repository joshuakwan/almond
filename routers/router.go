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
	"github.com/joshuakwan/almond/controllers"
)

func init() {
	namespaceRoot := "/api/v1"

	facadeNamespaceRoot := "/almond"
	nsFacade := beego.NewNamespace(namespaceRoot,
		beego.NSNamespace(facadeNamespaceRoot,
			beego.NSInclude(
				&controllers.AlmondController{},
			),
		),
	)
	beego.AddNamespace(nsFacade)
}
