package engines

import (
	"github.com/ahmetcanozcan/reves/sockets"
)

//Player :
type Player struct {
	id     string
	socket *sockets.Socket
}

//NewPlayer :
func NewPlayer(s *sockets.Socket) *Player {
	res := &Player{
		id:     s.GetID(),
		socket: s,
	}
	initialize(res)
	awakePlayer(res)
	return res
}

var awakePlayer func(*Player) = func(p *Player) {}

//OnPlayerInvoked :
func OnPlayerInvoked(f func(*Player)) {
	awakePlayer = f
}

//GetSocket :
func (p *Player) GetSocket() *sockets.Socket {
	return p.socket
}

//GetID :
func (p *Player) GetID() string {
	return p.id
}

func initialize(p *Player) {
	//TODO: Handle Key Events
}
