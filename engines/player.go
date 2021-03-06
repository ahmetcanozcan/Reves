package engines

import (
	"github.com/ahmetcanozcan/reves/sockets"
	"github.com/ahmetcanozcan/reves/sockets/messages"
)

//Player :
type Player struct {
	id         string
	socket     *sockets.Socket
	tempEvents []string
}

//NewPlayer :
func NewPlayer(s *sockets.Socket) *Player {
	res := &Player{
		id:         s.GetID(),
		socket:     s,
		tempEvents: make([]string, 0),
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

//Destroy :
func (p *Player) Destroy() {

	//Clear all temporary events
	for _, key := range p.tempEvents {
		p.socket.RemoveEvent(key)
	}

}

//On :
func (p *Player) On(name string, handler func(messages.Payload)) {
	p.tempEvents  = append(p.tempEvents, name)
	p.socket.On(name, handler)
}

//Emit :
func (p *Player) Emit(name string, payload messages.Payload) {
	p.socket.Emit(name, payload)
}

func initialize(p *Player) {
	//TODO: Handle Key Events
}
