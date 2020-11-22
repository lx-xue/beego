// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"myapp-beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsController{},
			),
			beego.NSRouter("/get-one/:id",&controllers.GoodsController{},"get:GetOne"),
		),

		beego.NSNamespace("/person",
			beego.NSInclude(
				&controllers.PersonController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
