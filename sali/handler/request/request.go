package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sali/handler/request/mutation"
	"sali/handler/request/query"
	"strconv"
	"strings"
)

type request interface {
	Parse(map[string]reflect.Type)
	ServeJSON(http.ResponseWriter)
}

func HandleRequest(httpres http.ResponseWriter, httpreq *http.Request, entities map[string]reflect.Type) {
	var req request
	switch strings.Split(httpreq.Header.Get("Content-Type"), ";")[0] {
	case "application/x-www-form-urlencoded":
		qu := query.Query{
			Method: httpreq.Method,
			Entity: httpreq.FormValue("entity"),
			Limit:  1,
			Offset: 0,
			Order:  "ID",
		}
		for id, v := range []string{"limit", "offset"} {
			if number := httpreq.FormValue(v); number != "" {
				u, err := strconv.ParseInt(number, 10, 64)
				if err != nil {
					qu.Errors = append(qu.Errors, err.Error())
				} else {
					switch id {
					case 0:
						qu.Limit = u
					case 1:
						qu.Offset = u
					}
				}
			}
		}
		if fields := httpreq.FormValue("fields"); fields != "" {
			fmt.Println("fields: ", fields)
			qu.Fields = strings.Split(strings.ReplaceAll(fields, " ", ""), ",")
		}
		if order := httpreq.FormValue("order"); order != "" {
			fmt.Println("order: ", order)
			qu.Order = order
		}
		if err := httpreq.ParseForm(); err != nil {
			qu.Errors = append(qu.Errors, err.Error())
		}
		for _, v := range []string{"entity", "fields", "orders", "limit", "order", "offset"} {
			delete(httpreq.Form, v)
		}
		qu.Filter = httpreq.Form
		req = &qu
	case "application/json":
		mu := mutation.Mutation{
			Method: httpreq.Method,
		}
		decoder := json.NewDecoder(httpreq.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&mu); err != nil {
			mu.Errors = append(mu.Errors, err.Error())
		}
		fmt.Println(mu)
		req = &mu
	}
	req.Parse(entities)
	req.ServeJSON(httpres)
}
