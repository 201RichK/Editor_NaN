package main

import (
	"log"
	"net/http"
	"sali/entities"
	"sali/handler"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("server listens on port 1234")
	h := handler.NewHandler([]interface{}{new(entities.Student), new(entities.Language), new(entities.Challenge), new(entities.Exercise), new(entities.Info), new(entities.Message), new(entities.Organisation), new(entities.Score), new(entities.Team), new(entities.Run)})
	http.ListenAndServe(":1234", h)
}
