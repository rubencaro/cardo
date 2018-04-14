package web

import (
	"fmt"
	"log"
	"net/http"

	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/websocket"
	"github.com/rubencaro/cardo/lib/web/sockets"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// to allow development websockets, to be configured differently on production
	CheckOrigin: func(r *http.Request) bool { return true },
}

// SocketsHandler handles the connection and dispatching
// of events to the client
func SocketsHandler(coll driver.Collection) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println("Socket connected")

		// go generateLogs(conn)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			err = sockets.Dispatch(conn, coll, string(msg))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// func generateLogs(conn *websocket.Conn) {
// 	for {
// 		err := send(conn, "bump")
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		time.Sleep(time.Second * 5)
// 		log.Println("Bump sent")
// 	}
// }
