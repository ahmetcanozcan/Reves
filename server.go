package reves

import (
	"fmt"
	"log"
	"net"

	"github.com/ahmetcanozcan/reves/sockets"
)

//Platform :
type Platform int32

const (
	//BROWSER :
	BROWSER Platform = 1
	//PC :
	PC Platform = 2
	//MOBILE :
	MOBILE Platform = 3
)

//ConfigurationStruct :
type ConfigurationStruct struct {
	HOST                    string
	PORT                    string
	IsAuthActive            bool
	GamePlatform            Platform
	BrowserStaticFolderPath string
	WebSocketConnectionURI  string
}

//Config :
var Config *ConfigurationStruct = &ConfigurationStruct{
	HOST:                    "localhost",
	PORT:                    "8080",
	IsAuthActive:            false,
	GamePlatform:            PC,
	BrowserStaticFolderPath: "public",
	WebSocketConnectionURI:  "/ws",
}

var cb func(*sockets.Socket)

//Start :
func Start() {
	if Config.GamePlatform == BROWSER {
		startBrowserServer()
	} else {
		startDefaultServer()
	}
}

func startDefaultServer() {
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
				wr := sockets.NewTCPWriterReader(&conn)
				socket := sockets.NewSocket(wr)
				cb(socket)
			}()
		}
	}
}

//HandleNewSocket :
func HandleNewSocket(callback func(socket *sockets.Socket)) {
	cb = callback
}
