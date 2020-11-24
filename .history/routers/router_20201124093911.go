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
	beego.Router("/default/index", &controllers.DefaultController{}, "get:Index") //参数id
	beego.Router("/default/adduser", &controllers.DefaultController{}, "post:AddUser")
	
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSRouter("/add",&controllers.UserController{},"post:Add"),
			beego.NSRouter("/getone",&controllers.UserController{},"get:GetOne"),
			beego.NSRouter("/GetAll",&controllers.UserController{},"get:GetAll"),
			beego.NSRouter("/put",&controllers.UserController{},"put:Put"),
			beego.NSRouter("/delete/:id",&controllers.UserController{},"get:Delete"),
		),

		beego.NSNamespace("/goods",
			// beego.NSInclude(
			// 	&controllers.GoodsController{},
			// ),
			beego.NSRouter("/post",&controllers.GoodsController{},"Post:Post"),
			beego.NSRouter("/getone/:id",&controllers.GoodsController{},"get:GetOne"),
		),

		beego.NSNamespace("/person",
			// beego.NSInclude(
			// 	&controllers.PersonController{},
			// ),
		),
	)
	beego.AddNamespace(ns)
}
