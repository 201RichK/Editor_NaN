package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/201RichK/Editor_NaN/models"
	"github.com/201RichK/Editor_NaN/utils"
	"net/http"
	"runtime"

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
	exercice, err := models.GetExerciceById(3)
	if err != nil {
		panic(err)
	}
	this.TplName = "index.html"
	this.Data["exercice"] = exercice
	this.Render()
}

func (this *MainController) Send() {
	logrus.Info(runtime.NumGoroutine())

	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Authorization, Content-Type")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if this.Ctx.Request.Method == http.MethodOptions {
		this.Abort("204")
		this.ServeJSON()
		return
	}

	exercice := make(map[string]interface{})
	result := make(map[string]interface{})
	token := make(map[string]interface{})
	err := json.NewDecoder(this.Ctx.Request.Body).Decode(&exercice)
	programhead := "package main \n import \"fmt\" \n " + exercice["source_code"].(string) + "\n func main() { \n " + "fmt.Println(Somme(5, 4))" + " \n }"
	exercice["source_code"] = programhead

	if err != nil {
		panic(err)
	}
	exercice["language_id"] = 22
	b, err := json.Marshal(exercice)
	fmt.Println("info exercice ",string(b))
	reader := bytes.NewReader(b)
	res, err := utils.MakeRequest(http.MethodPost, "application/json", "", reader)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(res.Body).Decode(&token)
	fmt.Println("token ",token)
	res, err = utils.MakeRequest(http.MethodGet, "application/x-www-form-urlencoded", token["token"].(string), nil)
	json.NewDecoder(res.Body).Decode(&result)
	logrus.Info("result ",result)
	this.Data["json"] = result
	this.ServeJSON()
}
