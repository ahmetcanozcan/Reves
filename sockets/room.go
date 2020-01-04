package sockets

import "fmt"

//Room : is an abstraction to group sockets
type Room struct {
	Name    string
	Sockets SocketList
}

var rooms []*Room = make([]*Room, 0)

//GetRoom : returns a room matched by given name
func GetRoom(name string) *Room {
	for i, val := range rooms {
		if val.Name == name {
			fmt.Println("Room", name, "returned")
			return rooms[i]
		}
	}
	CreateRoom(name)
	return GetRoom(name)
}

//CreateRoom : creates a room with given name
func CreateRoom(name string) {
	r := Room{
		Name:    name,
		Sockets: make(SocketList, 0),
	}
	rooms = append(rooms, &r)
}
