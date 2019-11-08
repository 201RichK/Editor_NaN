package controllers

import (
	"Editor_NaN/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"net/http"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	this.TplName = "index.html"
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
