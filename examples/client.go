package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error")
	}
	ioin := bufio.NewReader(os.Stdin)

	fmt.Fprintf(conn, "Init;\n")
	fmt.Println("sent Init")
	ioin.ReadLine()
	fmt.Fprintf(conn, "Hello;from:other-side;asd:asdasd;\n")
	fmt.Println("sent Hello")
	fmt.Fprintf(conn, "MatchMaking;\n")
	fmt.Println("sent MatchMaking")
	bufio.NewReader(conn).ReadLine()
}
