package routers

import (
	"Editor_NaN/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
    beego.Router("/send", &controllers.MainController{}, "options,post:Send")
}
