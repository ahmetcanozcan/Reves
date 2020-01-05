package sockets

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"net"
	sparser "tcpgameserver/sockets/parser"
)

//Socket : an abstraction for handling tcp connection
type Socket struct {
	events      map[string]func(sparser.Payload)
	id          string
	conn        *net.Conn
	initialized bool
}

//On : Dispatch an event handler to event listener
func (s *Socket) On(name string, handler func(sparser.Payload)) {
	s.events[name] = handler
}

//Emit : Send payload to client
func (s *Socket) Emit(name string, payload sparser.Payload) {
	message := fmt.Sprintf("Name : %s ; %s", name, payload.String())
	w := bufio.NewWriter(*s.conn)
	w.WriteString(message)
}

//IsOk : returns the socket complete its pre-communicaion tasks
func (s *Socket) IsOk() bool {
	return s.initialized
}

//NewSocket : Constructor
func NewSocket(conn *net.Conn) *Socket {
	s := Socket{
		events:      make(map[string]func(sparser.Payload)),
		id:          generateSocketID(),
		conn:        conn,
		initialized: true,
	}
	go s.listen()
	return &s
}

func (s *Socket) initialize(roomName string) {
	room := GetRoom(roomName)
	room.Sockets.Push(s)
	s.initialized = true
}

func (s *Socket) listen() {
	reader := bufio.NewReader(*s.conn)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		name, args, err := sparser.ParseSocketPayload(text)
		if err != nil {
			continue
		}
		if name == "Init" {
			s.initialize(args["roomName"])
			continue
		}
		f, ok := s.events[name]
		if ok && s.IsOk() {
			f(args)
		} else {
			fmt.Println("Event name doesn't found")
		}
	}
}

func generateSocketID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
