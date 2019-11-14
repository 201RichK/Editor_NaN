package utilsConf

import (
	"github.com/astaxie/beego"
)

/*
	Formulation de l'url pour la requete sur l'API
*/
func BaseUrl(token map[string]interface{}) string {
	url := "https://api.judge0.com/submissions/"
	if t, ok := token["token"].(string); ok {
		url = url + t
	}
	return url + "?base64_encoded=false&wait=false"
}

/*
	Configuration du header
*/
func SetHeader(this *beego.Controller) {
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Authorization, Content-Type")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
}

//Checher les erreur
func CheckError(er error) {
	if er != nil {
		panic(er)
	}
}
