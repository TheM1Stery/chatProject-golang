package chat

import (
	"slices"
	"sync"
)

type MessageType string

const (
	CreateChat MessageType = "create"
	LeaveChat  MessageType = "leave"
	JoinChat   MessageType = "join"
)

type User struct {
	username string
}

// this struct will act as an actor handler(ChatBroker is an actor)
type Room struct {
	// temporary type, should be uuid
	id     string
	Broker *ChatBroker
	name   string
	users  []User
}

func NewRoom(name string) Room {
	broker := NewChat()
	users := make([]User, 0, 10)
	var id string
	go broker.Start()
	return Room{
		id,
		broker,
		name,
		users,
	}
}

type ThreadSafeRoomList struct {
	rooms []Room
	mutex sync.Mutex
}

func NewRoomList() *ThreadSafeRoomList {
	return &ThreadSafeRoomList{
		mutex: sync.Mutex{},
		rooms: make([]Room, 0),
	}
}

func (roomList *ThreadSafeRoomList) Add(room *Room) {
	roomList.mutex.Lock()
	defer roomList.mutex.Unlock()

	roomList.rooms = append(roomList.rooms, *room)
}

func (roomList *ThreadSafeRoomList) Get(rid string) *Room {
	roomList.mutex.Lock()
	defer roomList.mutex.Unlock()

	index := slices.IndexFunc(roomList.rooms, func(room Room) bool { return rid == room.id })

	if index == -1 {
		return nil
	}

	return &roomList.rooms[index]
}
