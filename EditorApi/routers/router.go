package routers

import (
	"Editor_NaN/EditorApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get,post:Home")
	beego.Router("/ws", &controllers.WebScoketController{}, "get,post:Handle")
	beego.Router("/play", &controllers.HomeController{}, "get:Play")
	beego.Router("/run", &controllers.ProgramController{}, "post,options:RunProgram")
	beego.Router("/login", &controllers.UserController{}, "post,options:Login")
}
