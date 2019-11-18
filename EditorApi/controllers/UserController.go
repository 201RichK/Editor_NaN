package controllers


import (
	"Editor_NaN/EditorApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	utils "Editor_NaN/EditorApi/utils/option"
	"net/http"
)

type UserController struct {
	beego.Controller
}


func (this *UserController) Login() {

	utils.SetHeader(&this.Controller)

	//Annuler la requete si on a une methode OPTION
	if this.Ctx.Request.Method == http.MethodOptions {
		this.Abort("204")
		this.ServeJSON()
		return
	}


	user := models.User{
		Email: this.GetString("email"),
		Password: this.GetString("password"),
	}

	o := orm.NewOrm()
	u := new(models.User)
	o.QueryTable("User").Filter("Email", user.Email).Filter("Password", user.Password).One(u)
	if u != nil {
		this.StartSession()
		this.SetSession("connected", u)
	} else {
		this.Data["json"] = struct { Error string} { Error: "Email and Password don't correspond"}
	}
	this.ServeJSON()
}