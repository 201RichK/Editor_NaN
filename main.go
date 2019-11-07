package main

import (
	"html/template"
	"net/http"
)


var TPL *template.Template

func init()  {
	TPL = template.Must(template.ParseGlob("template/*.html"))
}

func main (){
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}


func Index(w http.ResponseWriter, r *http.Request) {
	TPL.ExecuteTemplate(w, "index.html", nil)
}