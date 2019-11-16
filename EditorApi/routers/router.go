package routers

import (
	"Editor_NaN/EditorApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/run", &controllers.ProgramController{}, "post,options:RunProgram")
	beego.Router("/login", &controllers.UserController{}, "post,options:Login")
}
