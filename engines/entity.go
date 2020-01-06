package engines

import (
	"github.com/ahmetcanozcan/reves/sockets"
)

//Entity :
type Entity interface {
	Start()
	Check()
	Update(int64)
	Share(room *sockets.Room)
}

//EntityList :
type EntityList []Entity

//Push :
func (l *EntityList) Push(e Entity) {
	*l = append(*l, e)
}

//Remove :
func (l *EntityList) Remove(i int) {
	(*l)[i] = (*l)[len((*l))-1]
	(*l)[len((*l))-1] = nil
	(*l) = (*l)[:len((*l))-1]
}
