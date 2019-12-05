package handler

import (
	"net/http"
	"os"
	"reflect"
	"sali/handler/request"
	"sali/handler/request/orm"
)

type Handler struct {
	entities map[string]reflect.Type
}

func NewHandler(entities []interface{}) *Handler {
	h := Handler{
		entities: make(map[string]reflect.Type),
	}
	h.registerEntity(entities)
	connectdb()
	return &h
}

func (h *Handler) registerEntity(entities []interface{}) {
	mp := make(map[string]reflect.Type)
	orm.RegisterModel(entities)
	for _, entity := range entities {
		etype := reflect.TypeOf(entity)
		if etype.Kind() == reflect.Ptr {
			etype = etype.Elem()
		}
		if etype.Kind() != reflect.Struct {
			panic("the entities must be of type Struct")
		}
		mp[etype.Name()] = etype
	}
	h.entities = mp
}

func connectdb() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}
	err = orm.RegisterDataBase("default", "postgres", "postgres://postgres:@172.17.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	orm.NewLog(os.Stdout)
	orm.RunSyncdb("default", false, true)
}

func (h *Handler) ServeHTTP(httpres http.ResponseWriter, httpreq *http.Request) {
	request.HandleRequest(httpres, httpreq, h.entities)
}
