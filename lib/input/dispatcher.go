package input

import (
	"fmt"
	"strings"

	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/websocket"
)

type dispatchData struct {
	conn   *websocket.Conn
	coll   driver.Collection
	msg    string
	prefix string
}

var routes = make(map[string]func(*dispatchData) error)

// Dispatch reads incoming message from given Conn and dispatches it
// to the right function
func Dispatch(conn *websocket.Conn, coll driver.Collection, msg string) error {
	data := &dispatchData{conn, coll, msg, extractPrefix(msg)}

	handler, ok := routes[data.prefix]
	if !ok {
		conn.Close()
		return fmt.Errorf("Unexpected msg: %s", msg)
	}
	return handler(data)
}

func extractPrefix(msg string) string {
	return strings.SplitN(msg, ": ", 2)[0]
}
