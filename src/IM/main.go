package main

import (
	_ "IM/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
)

const (
	AppVer = "0.0.1"
)

func main() {
	logs.Info(beego.BConfig.AppName, AppVer)

	// in html use i18n as tag map func to i18n.Tr
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
