package controller

import (
	"Editor_NaN/conf"
	"net/http"
)

type mainController struct {}

func InitMainController () *mainController {
	return &mainController{}
}

func (Mc mainController)Index(w http.ResponseWriter, r *http.Request) {
	conf.TPL.ExecuteTemplate(w, "index.html", nil)
}