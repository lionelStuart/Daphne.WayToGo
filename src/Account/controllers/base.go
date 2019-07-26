package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/sessions"
)

var (
	CookieStore *sessions.CookieStore
)

func init() {
	sessionKey := "test"
	CookieStore = sessions.NewCookieStore([]byte(sessionKey))
	CookieStore.Options = &sessions.Options{
		MaxAge: 1800,
	}
}

type baseController struct {
	beego.Controller
}
