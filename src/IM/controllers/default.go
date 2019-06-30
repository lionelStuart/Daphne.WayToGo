package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
)

type MainController struct {
	beego.Controller
	i18n.Locale
}

func (c *MainController) Prepare() {
	logs.Info("prepare for main controller")

	c.Lang = ""
	al := c.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5]
		if i18n.IsExist(al) {
			c.Lang = al
		}
	}

	if len(c.Lang) == 0 {
		c.Lang = "en-US"
	}

	c.Data["Lang"] = c.Lang
}

func (c *MainController) Get() {
	logs.Debug("main controller get")
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	listId := ListAllRoom()
	if listId != nil && listId.Len() > 0 {
		logs.Info("room list size", listId.Len())
		var rtList []uint64
		for r := listId.Front(); r != nil; r = r.Next() {
			rtList = append(rtList, r.Value.(uint64))
		}
		c.Data["goods"] = rtList
	} else {
		logs.Info("room list empty for now on")
		c.Data["goods"] = []uint64{1, 2, 3}
	}

	c.TplName = "index.tpl"
}

func (c *MainController) Join() {
	uname := c.GetString("uname")
	roomId := c.GetString("roomId")

	logs.Debug("main controller join ", uname)

	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}
	redirectUrl := fmt.Sprintf("/ws?uname=%s&roomId=%s", uname, roomId)
	logs.Debug("redirect to ", redirectUrl)
	c.Redirect(redirectUrl, 302)
	return
}
