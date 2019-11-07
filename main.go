package main

import (
	"Editor_NaN/controller"
	"net/http"
)




func main (){

	mc := controller.InitMainController()

	http.HandleFunc("/", mc.Index)
	http.ListenAndServe(":8080", nil)
}

