package controller

import (
	"Editor_NaN/conf"
	"Editor_NaN/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	//"github.com/sirupsen/logrus"
)

type mainController struct {}

func InitMainController () *mainController {
	return &mainController{}
}

type Exercie struct {
	sourceCode string `json:"source_code"`
	languagId int `json:"language_id"`
}

func (Mc mainController)Index(w http.ResponseWriter, r *http.Request) {
	conf.TPL.ExecuteTemplate(w, "index.html", nil)
}

func (Mc mainController) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers","X-Requested-With, Content-Type, Origin, Cache-Control, Pragma, Authorization, Accept, Accept-Encoding,Access-Control-Allow-Origin")
	w.Header().Add("Content-Type", "application/json")
	exercice := make(map[string]interface{})
	token := make(map[string]interface{})
	result := make(map[string]interface{})
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&exercice)
		if err != nil {
			panic(err)
		}
		exercice["language_id"] = 22
		b, err := json.Marshal(exercice)
		fmt.Println(string(b))
		reader := bytes.NewReader(b)
		res, err := utils.MakeRequest(http.MethodPost, "application/json", "", reader)
		json.NewDecoder(res.Body).Decode(&token)
		fmt.Println(token["token"].(string))
		time.Sleep(3 * time.Second)
		res, err = utils.MakeRequest(http.MethodGet, "application/x-www-form-urlencoded", token["token"].(string), nil)
		json.NewDecoder(res.Body).Decode(&result)
		logrus.Info(result)
	}

	m := make(map[string]interface{})
	m["stdout"] = result["stdout"]
	m["stderr"] = result["stderr"]
	m["time"] = result["time"]

	json.NewEncoder(w).Encode(m)
}



