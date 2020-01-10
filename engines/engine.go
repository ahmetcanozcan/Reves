package engines

import (
	"time"

	"github.com/ahmetcanozcan/reves/sockets"
)

//Engine :
type Engine struct {
	players   PlayerList
	entities  EntityList
	UR        int
	room      *sockets.Room
	onDestroy func()
}

//NewEngine :
func NewEngine(room *sockets.Room) *Engine {
	room.Type = sockets.GAME
	e := &Engine{UR: 20, entities: make(EntityList, 0), room: room, onDestroy: func() {}}
	e.players = make(PlayerList, 0)
	for _, socket := range room.Sockets {
		p := NewPlayer(socket)
		e.players.Push(p)
	}
	return e
}

//AddEntity :
func (e *Engine) AddEntity(entity Entity) {
	e.entities.Push(entity)
}

//OnDestroy :
func (e *Engine) OnDestroy(f func()) {
	e.onDestroy = f
}

func (e *Engine) check() bool {
	for _, entity := range e.entities {
		entity.Check()
	}
	if e.room.IsInactive() {
		return false
	}
	return true
}

func (e *Engine) update(deltaTime int64) {
	for _, entity := range e.entities {
		entity.Update(deltaTime)
	}
}

func (e *Engine) share() {
	for _, entity := range e.entities {
		entity.Share(e.players)
	}
}

func (e *Engine) destroy() {
	e.onDestroy()

}

//Start :
func (e *Engine) Start() {
	go func() {
		for _, entitiy := range e.entities {
			(entitiy).Start()
		}
		start := time.Now()
		for {
			time.Sleep(time.Duration(1000/e.UR) * time.Millisecond)
			if !e.check() {
				break
			}
			delta := time.Since(start)
			e.update(delta.Milliseconds())
			e.share()
			start = time.Now()
		}

		e.destroy()
	}()
}
