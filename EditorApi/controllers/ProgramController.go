package controllers

import (
	"Editor_NaN/EditorApi/models"
	remote "Editor_NaN/EditorApi/utils/Remote"
	utils "Editor_NaN/EditorApi/utils/option"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

type ProgramController struct {
	beego.Controller
}



/*

	RunProgram fait la requete sur l'api pour complier le code lorsque l'utilisateur fait Run

*/

func (this *ProgramController) RunProgram() {
	utils.SetHeader(this.Controller)
	if this.Ctx.Request.Method == http.MethodOptions {
		this.Abort("204")
		this.ServeJSON()
		return
	}
	exerciceModel := new(models.ExerciceModel)
	json.NewDecoder(this.Ctx.Request.Body).Decode(&exerciceModel.Program)
	programHeader := "package main \nimport \"fmt\"  \n" + exerciceModel.Program["source_code"].(string) + "\nfunc main() { \n\n " + "fmt.Println(somme(5, 4))" + "\n\n }"
	exerciceModel.Program["source_code"] = programHeader
	exerciceModel.Program["language_id"] = 22
	rmt := remote.NewRemote(remote.RemoteConfig{})
	res, err := rmt.POST(remote.RequestConfig{
		URL:  utils.BaseUrl(exerciceModel.Token),
		Body: exerciceModel.Program,
	})
	utils.CheckError(err)
	json.NewDecoder(res.Body).Decode(&exerciceModel.Token)
	res, err = rmt.GET(remote.RequestConfig{
		URL: utils.BaseUrl(exerciceModel.Token),
	})
	utils.CheckError(err)
	json.NewDecoder(res.Body).Decode(&exerciceModel.Result)
	fmt.Println(exerciceModel.Result)
	this.ServeJSON()
}

