package controllers

import (
	. "IM/models"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WebSocketController struct {
	MainController
}

func (c *WebSocketController) Get() {
	logs.Info("websocket controller get")

	uname := c.GetString("uname")
	roomId := c.GetString("roomId")

	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	c.TplName = "websocket.html"
	c.Data["IsWebSocket"] = true
	c.Data["UserName"] = uname
	c.Data["RoomId"] = roomId
}

func (c *WebSocketController) Join() {
	uname := c.GetString("uname")
	roomId, _ := c.GetUint64("roomid")

	logs.Info("websocket controller join ", uname, roomId)

	if len(uname) == 0 {
		c.Redirect("/", 302)
		return
	}

	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "not a websocket handshake", 400)
	} else if err != nil {
		logs.Error("can not setup socket connection ", err)
	}

	ch := make(chan interface{}, 1)

	userInfo := UserInfo{
		RoomId: roomId,
		Uname:  uname,
	}

	user := User{
		UserInfo: userInfo,
		Ws:       ws,
		Chan:     &ch,
	}
	defer LeaveRoom(user)

	logs.Info("now join people ", user.Uname, user.RoomId)

	subRoomId := JoinRoom(user)
	if subRoomId == 0 {
		logs.Error("fail join room for", user.Uname)
		return
	}
	user.RoomId = subRoomId
	userInfo.RoomId = subRoomId
	PostMessage(Event{
		Type:      EVENT_JOIN,
		UserInfo:  userInfo,
		Timestamp: int(time.Now().Unix()),
		Content:   "user join",
	})

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}

		PostMessage(Event{
			Type:      EVENT_MESSAGE,
			UserInfo:  userInfo,
			Timestamp: int(time.Now().Unix()),
			Content:   string(p),
		})
	}

}
