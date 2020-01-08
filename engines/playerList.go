package engines

//PlayerList : a list implementation for sockets
type PlayerList []*Player

//Push : Add a socket to end of the list
func (l *PlayerList) Push(s *Player) {
	*l = append(*l, s)
}

//Remove : Remove a socket from given index
func (l *PlayerList) Remove(i int) {
	(*l)[i] = (*l)[len((*l))-1]
	(*l)[len((*l))-1] = &Player{}
	(*l) = (*l)[:len((*l))-1]
}
