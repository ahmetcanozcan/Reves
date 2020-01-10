package sockets

import "fmt"

//SocketList : a list implementation for sockets
type SocketList []*Socket

//Push : Add a socket to end of the list
func (l *SocketList) Push(s *Socket) {
	*l = append(*l, s)
}

//Remove : Remove a socket from given index
func (l *SocketList) Remove(i int) {
	(*l)[i] = (*l)[len((*l))-1]
	(*l)[len((*l))-1] = &Socket{}
	(*l) = (*l)[:len((*l))-1]
}

//RemoveElement : Remove a socket from given index
func (l *SocketList) RemoveElement(el *Socket) {
	var i int
	for ind, sck := range *l {
		if sck.Equals(el) {
			i = ind
			break
		}
	}
	(*l)[i] = (*l)[len((*l))-1]
	(*l)[len((*l))-1] = &Socket{}
	(*l) = (*l)[:len((*l))-1]
}

//Print : Display all sockets in the list
func (l SocketList) Print() {
	fmt.Println()
	for _, val := range l {
		fmt.Println(val.id)
	}
}
