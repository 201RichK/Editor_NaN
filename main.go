package main

import (
	_ "github.com/201RichK/workspace/Editor_NaN/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	//run the Db on sync mode
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}

	//enable session
	beego.BConfig.WebConfig.Session.SessionOn = true

	//run the app
	beego.Run()
}
