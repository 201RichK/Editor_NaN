package controllers

import (
	"fmt"
	"net/http"
	"workspace/Editor_NaN/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	this.StartSession()
	v := this.GetSession("connected")
	if v == nil {
		this.Ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
	exercices, err := models.GetExerciceById(1)
	if err != nil {
		panic(err)
	}
	this.TplName = "index.html"
	fmt.Println(exercices)
	this.Data["message"] = exercices
	this.Render()
}
