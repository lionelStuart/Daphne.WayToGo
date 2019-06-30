package controllers

import (
	. "IM/models"
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"sync/atomic"
)

//edit from here

type ChatRoom struct {
	roomId      uint64
	subscribers list.List
	waitingList list.List
}

var (
	roomManager        = NewRoomManager()
	subscribeChannel   = make(chan User, 10)
	unsubscribeChannel = make(chan User, 10)
	publishChannel     = make(chan Event, 10)
)

func JoinRoom(user User) uint64 {
	select {
	case subscribeChannel <- user:
		return (<-*user.Chan).(uint64)
	}
}

func LeaveRoom(user User) {
	select {
	case unsubscribeChannel <- user:
	default:
		fmt.Println("failure call leave")
	}
}

func PostMessage(event Event) {
	select {
	case publishChannel <- event:
	}
}

func ListAllRoom() *list.List {
	return roomManager.ListRoom()
}

type RoomManager struct {
	roomIdInc uint64
	rooms     map[uint64]*ChatRoom
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomIdInc: 0,
		rooms:     make(map[uint64]*ChatRoom),
	}
}

func (m *RoomManager) CreateRoom() *ChatRoom {
	atomic.AddUint64(&m.roomIdInc, 1)
	m.rooms[m.roomIdInc] = &ChatRoom{
		roomId: m.roomIdInc,
	}
	logs.Info("create room id", m.roomIdInc)
	return m.rooms[m.roomIdInc]
}

func (m *RoomManager) DestroyRoom(roomId uint64) {
	if m.ExistsRoom(roomId) {
		delete(m.rooms, roomId)
	}
}

func (m *RoomManager) ExistsRoom(roomId uint64) bool {
	_, ok := m.rooms[roomId]
	return ok
}

func (m *RoomManager) ListRoom() *list.List {
	l := list.New()
	for room := range m.rooms {
		l.PushBack(room)
	}
	return l
}

func (m *RoomManager) GetRoom(roomId uint64) (*ChatRoom, bool) {
	if m.ExistsRoom(roomId) {
		return m.rooms[roomId], true
	}
	return nil, false
}

func (r *ChatRoom) Join(user *User) {
	if !r.HasUser(user.Uname) {
		r.subscribers.PushBack(user)
		logs.Info("chat room join", user.Uname)
	}
}

func (r *ChatRoom) Leave(uname string) {
	for user := r.subscribers.Front(); user != nil; user = user.Next() {
		if user.Value.(*User).Uname == uname {
			ws := user.Value.(*User).Ws
			if ws != nil {
				ws.Close()
			}
			r.subscribers.Remove(user)
		}
	}
}

func (r *ChatRoom) Broadcast(event *Event) {
	logs.Info("unhandled broadcast")

	data, err := json.Marshal(event)
	if err != nil {
		logs.Error("fail to marshal event ", err)
		return
	}
	msg := Event{}
	err = json.Unmarshal(data, &msg)
	logs.Info("have msg ", msg.Content)

	for sub := r.subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(*User).Ws
		if ws != nil {
			logs.Info("send message to ", sub.Value.(*User).Uname)
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				unsubscribeChannel <- *sub.Value.(*User)
			}
		}
	}

}

func (r *ChatRoom) HasUser(uname string) bool {
	for user := r.subscribers.Front(); user != nil; user = user.Next() {
		if user.Value.(*User).Uname == uname {
			return true
		}
	}
	return false
}

func (r *ChatRoom) UserCount() int {
	return r.subscribers.Len()
}

//edit end here

// This function handles all incoming chan messages.
func chatLoop() {
	for {
		select {
		case sub := <-subscribeChannel:
			if sub.RoomId == 0 {
				//新创建
				newRoom := roomManager.CreateRoom()
				newRoom.Join(&sub)
				logs.Info("create room ", newRoom.roomId)
				*sub.Chan <- newRoom.roomId
			} else if room, ok := roomManager.GetRoom(sub.RoomId); ok {
				//加入
				room.Join(&sub)
				logs.Info("join room ", room.roomId)
				*sub.Chan <- room.roomId
			} else {
				logs.Error("room id not exists or create failure ", sub.RoomId)
			}
		case unsub := <-unsubscribeChannel:
			if room, ok := roomManager.GetRoom(unsub.RoomId); ok {
				room.Leave(unsub.Uname)
				if room.UserCount() == 0 {
					//如果人数为空应当销毁房间
					roomManager.DestroyRoom(room.roomId)
				}
			}
		case pub := <-publishChannel:
			logs.Info("do you broadcast")
			if room, ok := roomManager.GetRoom(pub.UserInfo.RoomId); ok {
				logs.Info("pub room number ", room.roomId)
				room.Broadcast(&pub)
			} else {
				logs.Info("fail get room ", pub.UserInfo.RoomId)
			}
		}

	}
}

func init() {
	go chatLoop()
}
