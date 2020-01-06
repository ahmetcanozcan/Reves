package reves

import (
	"fmt"
	"log"
	"net"
	"reves/sockets"
)

//Server configuration constans
const (
	HOST string = "localhost"
	PORT string = "8080"
)

//ListenSocket :
func ListenSocket(cb func(socket *sockets.Socket)) {
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
			go func() {
				socket := sockets.NewSocket(&conn)
				cb(socket)
			}()
		}
	}
}
