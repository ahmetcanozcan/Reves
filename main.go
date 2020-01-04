package main

import (
	"fmt"
	"log"
	"net"
	"tcpgameserver/sockets"
	sparser "tcpgameserver/sockets/parser"
)

//Server configuration constans
const (
	HOST string = "localhost"
	PORT string = "8080"
)

func main() {
	//creating listener for handling incomming requests
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Failed to listen port", PORT)
	}
	fmt.Println("Listening port", PORT, "...")
	//Close listener when app closed
	defer listener.Close()
	for {
		//Accept incomming request
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept  a request.")
			continue
		} else {
			//Handle connection in another thread
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	socket := sockets.NewSocket(&conn)
	socket.On("Hello", func(payload sparser.Payload) {
		for k, v := range payload {
			fmt.Println(k, v)
		}
	})
}
