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

	RunProgram fait la requete sur l'api pour compiler le code et renvoyer un resultat lorsque l'utilisateur fait un Run

*/
func (this *ProgramController) RunProgram() {

	utils.SetHeader(&this.Controller)

	//Annuler la requete si on a une methode OPTION
	if this.Ctx.Request.Method == http.MethodOptions {
		this.Abort("204")
		this.ServeJSON()
		return
	}

	//models de l'exercice
	exerciceModel := new(models.ExerciceModel)

	//Recevoir l'exercice venu de la page de composition
	json.NewDecoder(this.Ctx.Request.Body).Decode(&exerciceModel.Program)

	programHeader := "package main \nimport \"fmt\"  \n" + exerciceModel.Program["source_code"].(string) + "\nfunc main() { \n\n " + "fmt.Println(somme(5, 4))" + "\n\n }"
	exerciceModel.Program["source_code"] = programHeader
	exerciceModel.Program["language_id"] = 22
	fmt.Println(exerciceModel.Program)

	//poster l'exercice ver l'API
	rmt := remote.NewRemote(remote.RemoteConfig{})
	res, err := rmt.POST(remote.RequestConfig{
		URL:  utils.BaseUrl(exerciceModel.Token),
		Body: exerciceModel.Program,
	})
	utils.CheckError(err)

	//Decoder la respons puis faire un GET sur l'API pour le resultat
	json.NewDecoder(res.Body).Decode(&exerciceModel.Token)
	res, err = rmt.GET(remote.RequestConfig{
		URL: utils.BaseUrl(exerciceModel.Token),
	})
	utils.CheckError(err)

	//Decoder la respons puis recuperer  le resultat
	json.NewDecoder(res.Body).Decode(&exerciceModel.Result)
	fmt.Println(exerciceModel.Result)
	m := make(map[string]interface{})

	//Envoyer la reponse a la vue react
	m["time"] = exerciceModel.Result["time"]
	m["stdout"] = exerciceModel.Result["stdout"]
	m["stderr"] = exerciceModel.Result["stderr"]
	this.Data["json"] = m
	this.ServeJSON()
}