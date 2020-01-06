package main

import (
	"fmt"
	"reves"
	"reves/engines"
	"reves/engines/players"
	"reves/sockets"
	"reves/sockets/messages"
)

func main() {

	reves.Config.PORT = "3001"

	sockets.WhenMatchmakingRoomIsFilled(func(r *sockets.Room) {
		engine := engines.NewEngine(r)
		for _, socket := range r.Sockets {
			p := players.NewPlayer(socket)
			engine.AddEntity(p)
		}
		engine.Start()
	})

	reves.HandleNewSocket(func(socket *sockets.Socket) {

		socket.On("Hello", func(payload messages.Payload) {
			for k, v := range payload {
				fmt.Println(k, v)
			}
		})

	})

	reves.Start()

}
