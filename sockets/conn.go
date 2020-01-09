package sockets

import (
	"bufio"
	"fmt"
	"net"

	"github.com/gorilla/websocket"
)

//WriterReader :
type WriterReader interface {
	Read() (string, error)
	Write(string) bool
}

//TCPWriterReader :
type TCPWriterReader struct {
	conn   *net.Conn
	reader *bufio.Reader
}

//NewTCPWriterReader :
func NewTCPWriterReader(conn *net.Conn) *TCPWriterReader {
	r := TCPWriterReader{
		conn:   conn,
		reader: bufio.NewReader(*conn),
	}
	return &r
}

//Write :
func (t *TCPWriterReader) Write(message string) bool {
	fmt.Fprintf(*t.conn, message)
	return true
}

//Read :
func (t *TCPWriterReader) Read() (string, error) {
	text, err := t.reader.ReadString('\n')
	return text, err
}

//WebSocketWriterReader :
type WebSocketWriterReader struct {
	conn *websocket.Conn
}

//NewWebSocketWriterReader :
func NewWebSocketWriterReader(conn *websocket.Conn) *WebSocketWriterReader {
	r := WebSocketWriterReader{conn: conn}
	return &r
}

//Write :
func (wr *WebSocketWriterReader) Write(message string) bool {
	err := wr.conn.WriteMessage(websocket.TextMessage, []byte(message))
	return err == nil
}

//Read :
func (wr *WebSocketWriterReader) Read() (string, error) {
	_, text, err := wr.conn.ReadMessage()
	return string(text), err
}
