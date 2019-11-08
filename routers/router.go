package routers

import (
	"workspace/Editor_NaN/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/login", &controllers.UserController{}, "get:LoginPage;post:Login")

	beego.Router("/send", &controllers.MainController{}, "options,post:Send")
}
