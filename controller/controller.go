package controller

import (
	"Editor_NaN/conf"
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
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
	var exercie, result =  map[string]interface{}
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(exercie)
		if err != nil {
			panic(err)
		}
		exercie = 22
		b, err := json.Marshal(exercie)
		fmt.Println(string(b))
		reader := bytes.NewReader(b)
		client := new(http.Client)
		request, err := http.NewRequest("POST", "https://api.judge0.com/submissions/?base64_encoded=false&wait=false", reader)
		request.Header.Set("Content-Type", "application/json")
		res, err := client.Do(request)
		if err != nil {
			panic(err)
		}

	}

	json.NewEncoder(w).Encode("ok")
	logrus.Info(result)
}

