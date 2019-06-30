package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"strings"
)

var langTypes []string // Languages that are supported.

func init() {
	// 从app.conf 获取 已经注册的语言类型名
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// 从conf目录下获取注册的语言ini 并配置到i18n
	for _, lang := range langTypes {
		logs.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			logs.Error("Fail to set message file:", err)
			return
		}
	}
}
