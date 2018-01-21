package web

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	driver "github.com/arangodb/go-driver"
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
func SocketsHandler(coll driver.Collection) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println("Socket connected")

		go generateLogs(conn)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			err = dispatch(conn, coll, string(msg))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

var logRE = regexp.MustCompile(`^log: `)

func dispatch(conn *websocket.Conn, coll driver.Collection, msg string) error {
	switch {
	case logRE.MatchString(msg):
		dispatchLog(coll, msg)
		return nil
	default:
		conn.Close()
		return fmt.Errorf("Unexpected msg: %s", msg)
	}
}

func dispatchLog(coll driver.Collection, msg string) {
	meta, err := coll.CreateDocument(nil, map[string]string{"msg": msg})
	if err != nil {
		fmt.Println("failed to create document: ", err, "\nmeta: ", meta)
	}
	log.Println("Created new document for ", msg)
}

func send(conn *websocket.Conn, msg string) error {
	return conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func generateLogs(conn *websocket.Conn) {
	for {
		err := send(conn, "bump")
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second * 5)
		log.Println("Bump sent")
	}
}
