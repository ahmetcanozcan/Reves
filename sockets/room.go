package sockets

//RoomType :
type RoomType int

const (
	//DEFAULT :
	DEFAULT RoomType = 0
	//MATCHMAKING :
	MATCHMAKING RoomType = 1
	//CHAT :
	CHAT RoomType = 2
	//GAME :
	GAME RoomType = 3
)

//Room : is an abstraction to group sockets
type Room struct {
	Name    string
	Sockets SocketList
	Type    RoomType
	limit   int
}

// AddSocket :
func (r *Room) AddSocket(s *Socket) bool {
	res := false
	if len(r.Sockets) < r.limit {
		r.Sockets.Push(s)
		res = true
	}
	if len(r.Sockets) == r.limit {
		if r.Type == MATCHMAKING {
			whenMatchmakingRoomIsFilled(r)
		}

		whenAnyRoomIsFilled(r)
	}
	return res
}

var emptyRoomFunction func(*Room) = func(r *Room) {}
var whenMatchmakingRoomIsFilled func(*Room) = emptyRoomFunction
var whenAnyRoomIsFilled func(*Room) = emptyRoomFunction

//WhenAnyRoomIsFilled :
func WhenAnyRoomIsFilled(f func(*Room)) {
	whenAnyRoomIsFilled = f
}

//WhenMatchmakingRoomIsFilled :
func WhenMatchmakingRoomIsFilled(f func(*Room)) {
	whenMatchmakingRoomIsFilled = f
}

var rooms []*Room = make([]*Room, 0)

//GetRoom : returns a room matched by given name
func GetRoom(name string) *Room {
	for i, val := range rooms {
		if val.Name == name {
			return rooms[i]
		}
	}
	CreateRoom(name, DEFAULT)
	return GetRoom(name)
}

//GetMatchMakingRoom :
func GetMatchMakingRoom() *Room {
	for i, val := range rooms {
		if val.Type == MATCHMAKING {
			return rooms[i]
		}
	}
	return nil
}

//CreateRoom : creates a room with given name
func CreateRoom(name string, rtype RoomType) {
	r := Room{
		Name:    name,
		Sockets: make(SocketList, 0),
	}
	rooms = append(rooms, &r)
}
