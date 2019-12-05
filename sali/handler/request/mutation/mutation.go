package mutation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Mutation struct {
	Method   string
	Errors   []string
	Entity   string
	Fields   interface{}
	Filter   map[string]interface{}
	Rel string
	NM2M     string
	NReverse string
	M2M      interface{}
	Reverse  interface{}
}

func (mu *Mutation) POST(httpres http.ResponseWriter) int64 {
	if id, err := mu.create(); err != nil {
		httpres.WriteHeader(301)
		fmt.Println(err)
		return 0
	} else {
		httpres.WriteHeader(200)
		return id
	}
}

func (mu *Mutation) PUT(httpres http.ResponseWriter) {
	if _, err := mu.update(); err != nil {
		fmt.Println(err)
		httpres.WriteHeader(301)
	} else {
		httpres.WriteHeader(200)
	}
}

func (mu *Mutation) PATCH(httpres http.ResponseWriter) {
	if _, err := mu.update(); err != nil {
		httpres.WriteHeader(301)
	} else {
		httpres.WriteHeader(200)
	}
}

func (mu *Mutation) DELETE(httpres http.ResponseWriter) {
	if _, err := mu.delete(); err != nil {
		httpres.WriteHeader(301)
	} else {
		httpres.WriteHeader(200)
	}
}

func (mu *Mutation) ServeJSON(httpres http.ResponseWriter) {
	response := struct {
		Operation string      `json:"operation"`
		Errors    []string    `json:"errors,omitempty"`
		Data      interface{} `json:"data"`
	}{
		Operation: "mutation",
		Errors:    mu.Errors,
	}
	if len(response.Errors) == 0 {
		switch mu.Method {
		case "POST":
			id := mu.POST(httpres)
			response.Data = map[string]interface{}{
				"ID": id,
			}
		case "PUT":
			mu.PUT(httpres)
		case "PATCH":
			mu.PATCH(httpres)
		case "DELETE":
			mu.DELETE(httpres)
		default:
		}
	}
	if err := json.NewEncoder(httpres).Encode(response); err != nil {
		log.Println(err)
	}
}
