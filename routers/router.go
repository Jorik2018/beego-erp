// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-erp/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	userController := &controllers.UserController{}
	carsController := &controllers.CarsControllers{}

	user := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/login", userController, "post:Login"),
			// beego.NSRouter("/get_all", userController, "post:GetAll"),
			beego.NSRouter("/add_user", userController, "post:RegisterUser"),
			// beego.NSRouter("/get_perticular_user", userController, "post:Get"),
			// beego.NSRouter("/delete_user", userController, "post:Delete"),
			beego.NSRouter("/login_user", userController, "post:LoginUser"),
		),
		beego.NSNamespace("/cars",
			beego.NSRouter("/register_car", carsController, "post:RegisterCar"),
			beego.NSRouter("/update_cars", carsController, "post:UpdateCar"),
			beego.NSRouter("/delete_cars", carsController, "post:DeleteCar"),
			beego.NSRouter("/fetch_cars", carsController, "post:FetchCar"),
		),
	)

	beego.AddNamespace(user)
}
