package controllers


import (
	"Editor_NaN/EditorApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}


func (this *UserController) Login() {
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