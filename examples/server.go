package main

import (
	"fmt"
	"tcpgameserver"
	"tcpgameserver/sockets"
	sparser "tcpgameserver/sockets/parser"
)

func main() {

	tcpgameserver.ListenSocket(func(socket *sockets.Socket) {

		socket.On("Hello", func(payload sparser.Payload) {
			for k, v := range payload {
				fmt.Println(k, v)
			}
		})

	})

}
