package models

import (
	"container/list"
	"github.com/gorilla/websocket"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type UserInfo struct {
	RoomId uint64
	Uname  string
}

type User struct {
	UserInfo
	Ws   *websocket.Conn
	Chan *chan interface{}
}

type Event struct {
	Type      EventType // JOIN, LEAVE, MESSAGE
	UserInfo  UserInfo
	Timestamp int // Unix timestamp (secs)
	Content   string
}

const archiveSize = 20

// Event archives.
var archive = list.New()

// NewArchive saves new event to archive list.
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

// GetEvents returns all events after lastReceived.
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
