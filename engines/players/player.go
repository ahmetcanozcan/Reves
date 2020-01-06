package players

import (
	"reves/sockets"
	"reves/sockets/messages"
)

//Player :
type Player struct {
	id     string
	socket *sockets.Socket
}

//NewPlayer :
func NewPlayer(s *sockets.Socket) *Player {
	return &Player{
		id:     s.GetID(),
		socket: s,
	}
}

//Start :
func (p *Player) Start() {
	p.socket.On("KeyPress", func(payload messages.Payload) {

	})
}

//Check :
func (p *Player) Check() {

}

//Update :
func (p *Player) Update(deltaTime int64) {

}

//Share :
func (p *Player) Share(room *sockets.Room) {

}
