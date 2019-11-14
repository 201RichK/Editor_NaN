package main

import (
	_ "Editor_NaN/EditorApi/models"
	_ "Editor_NaN/EditorApi/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
