package query

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Query struct {
	Limit  int64
	Offset int64
	Entity string
	Method string
	Order  string
	Errors []string
	Fields []string
	Filter interface{}
}

func (qu *Query) ServeJSON(httpres http.ResponseWriter) {
	response := struct {
		Operation string      `json:"operation"`
		Errors    []string    `json:"errors,omitempty"`
		Data      interface{} `json:"data"`
	}{
		Operation: "query",
		Errors:    qu.Errors,
	}
	if len(response.Errors) == 0 {
		switch qu.Method {
		case "GET":
			if _, data, err := qu.read(); err != nil {
				httpres.WriteHeader(301)
			} else {
				httpres.WriteHeader(200)
				response.Data = data
			}
		case "HEAD":
			if id, _, err := qu.read(); err != nil {
				httpres.WriteHeader(301)
			} else {
				httpres.Header().Set("id", strconv.Itoa(int(id)))
				httpres.WriteHeader(301)
			}
		default:
		}
	}
	if err := json.NewEncoder(httpres).Encode(response); err != nil {
		log.Println(err)
	}
}
