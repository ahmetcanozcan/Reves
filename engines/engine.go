package engines

import (
	"time"

	"github.com/ahmetcanozcan/reves/sockets"
)

//Engine :
type Engine struct {
	room     *sockets.Room
	entities EntityList
	UR       int
}

//NewEngine :
func NewEngine(room *sockets.Room) *Engine {
	return &Engine{
		room: room,
	}
}

//AddEntity :
func (e *Engine) AddEntity(entity Entity) {
	e.entities.Push(entity)
}

func (e *Engine) check() {
	for _, entity := range e.entities {
		entity.Check()
	}
}

func (e *Engine) update(deltaTime int64) {
	for _, entity := range e.entities {
		(entity).Update(deltaTime)
	}
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
			e.check()
			delta := time.Since(start)
			e.update(delta.Milliseconds())
			start = time.Now()
		}
	}()
}
