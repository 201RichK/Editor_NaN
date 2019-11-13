package utilsConf

import (
	"github.com/astaxie/beego"
)


func BaseUrl(token map[string]interface{}) string {
	url := "https://api.judge0.com/submissions/"
	if t, ok := token["token"].(string); ok {
		url = url + t
	}
	return url + "?base64_encoded=false&wait=false"
} 

func SetHeader(this beego.Controller) {
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Authorization, Content-Type")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
}


func CheckError(er error) {
	if er != nil {
		panic(er)
	}
}

