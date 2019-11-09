package main

import (
	//"runtime"
	_ "workspace/Editor_NaN/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/sirupsen/logrus"
)

func main() {

	//run the Db on sync mode
	//logrus.Info(runtime.GOMAXPROCS(runtime.NumCPU()))
	//logrus.Info(runtime.NumGoroutine())
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}

	//enable session
	beego.BConfig.WebConfig.Session.SessionOn = true

	//run the app
	beego.Run()
}
