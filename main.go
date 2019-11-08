package main

import (
	_ "workspace/Editor_NaN/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
