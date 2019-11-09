package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/201RichK/workspace/Editor_NaN/models"
	"github.com/201RichK/workspace/Editor_NaN/utils"

	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
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

func (this *MainController) Send() {
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Authorization, Content-Type")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	logrus.Info(this.Ctx.Request.Method)
	if this.Ctx.Request.Method == http.MethodOptions {
		this.Abort("204")
		this.ServeJSON()
		return
	}

	exercice := make(map[string]interface{})
	token := make(map[string]interface{})
	result := make(map[string]interface{})

	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&exercice)
	if err != nil {
		panic(err)
	}
	exercice["language_id"] = 22
	b, err := json.Marshal(exercice)
	fmt.Println(string(b))
	reader := bytes.NewReader(b)
	res, err := utils.MakeRequest(http.MethodPost, "application/json", "", reader)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(res.Body).Decode(&token)
	fmt.Println(token)
	res, err = utils.MakeRequest(http.MethodGet, "application/x-www-form-urlencoded", token["token"].(string), nil)
	json.NewDecoder(res.Body).Decode(&result)
	logrus.Info(result)
	this.Data["json"] = result
	this.ServeJSON()
}
