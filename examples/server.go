package main

import (
	"fmt"
	"reves"
	"reves/sockets"
	"reves/sockets/messages"
)

func main() {

	reves.ListenSocket(func(socket *sockets.Socket) {

		socket.On("Hello", func(payload messages.Payload) {
			for k, v := range payload {
				fmt.Println(k, v)
			}
		})

	})

}
