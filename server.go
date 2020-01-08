package reves

import (
	"fmt"
	"log"
	"net"

	"github.com/ahmetcanozcan/reves/sockets"
)

//ConfigurationStruct :
type ConfigurationStruct struct {
	HOST         string
	PORT         string
	IsAuthActive bool
}

//Config :
var Config *ConfigurationStruct = &ConfigurationStruct{
	HOST:         "localhost",
	PORT:         "8080",
	IsAuthActive: false,
}

var cb func(*sockets.Socket)

//Start :
func Start() {
	//creating listener for handling incomming requests
	listener, err := net.Listen("tcp", Config.HOST+":"+Config.PORT)
	if err != nil {
		fmt.Println("Failed to listen port", Config.PORT)
	}
	fmt.Println("Listening port", Config.PORT, "...")
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

//HandleNewSocket :
func HandleNewSocket(callback func(socket *sockets.Socket)) {
	cb = callback
}
