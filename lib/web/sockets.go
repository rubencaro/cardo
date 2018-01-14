package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// to allow development websockets, to be configured differently on production
	CheckOrigin: func(r *http.Request) bool { return true },
}

// SocketsHandler handles the connection and dispatching
// of events to the client
func SocketsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Socket connected")

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		if string(msg) == "ping" {
			fmt.Println("ping")
			time.Sleep(2 * time.Second)
			err = conn.WriteMessage(msgType, []byte("pong"))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			conn.Close()
			fmt.Println(string(msg))
			return
		}
	}
}
