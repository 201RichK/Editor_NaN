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
	mux := http.NewServeMux()

	//Routes
	mux.HandleFunc("/", mc.Index)
	mux.HandleFunc("/send", mc.Send)




	//lancer le server
	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	log.Info("server on listen on http://localhost:"+port)
	http.ListenAndServe(fmt.Sprintf(port), mux)
}

