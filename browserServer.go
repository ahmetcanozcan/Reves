package reves

import (
	"net/http"

	"github.com/ahmetcanozcan/reves/sockets"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func startBrowserServer() {
	http.HandleFunc(Config.WebSocketConnectionURI, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		wr := sockets.NewWebSocketWriterReader(conn)
		socket := sockets.NewSocket(wr)

		cb(socket)

	})

	fs := http.FileServer(http.Dir(Config.BrowserStaticFolderPath))
	http.Handle("/", fs)

	http.ListenAndServe(":"+Config.PORT, nil)

}
