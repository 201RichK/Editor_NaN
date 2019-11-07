package main

import (
	"Editor_NaN/controller"
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"os"
)




func main (){

	//initialisation du controller
	mc := controller.InitMainController()

	//Routes
	http.HandleFunc("/", mc.Index)



	//lancer le server
	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	log.Info("server on listen on http://localhost:"+port)
	http.ListenAndServe(fmt.Sprintf(port), nil)
}

