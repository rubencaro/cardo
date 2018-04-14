package sockets

import (
	"encoding/json"
	"fmt"
	"strings"

	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/websocket"
	"github.com/rubencaro/cardo/lib/cards"
	"github.com/rubencaro/cardo/lib/types"
)

var routes = make(map[string]func(*types.DispatchData) error)

func init() {
	routes["cards_addCard"] = parse(cards.HandleAddCard)
}

// Dispatch reads incoming message from given Conn and dispatches it
// to the right function
func Dispatch(conn *websocket.Conn, coll driver.Collection, msg string) error {
	prefix, payload := splitMsg(msg)
	data := &types.DispatchData{conn, coll, prefix, payload, nil}

	handler, ok := routes[data.Prefix]
	if !ok {
		conn.Close()
		return fmt.Errorf("Unexpected msg: %s", msg)
	}
	return handler(data)
}

func splitMsg(msg string) (string, string) {
	arr := strings.SplitN(msg, ": ", 2)
	return arr[0], arr[1]
}

// parse is a middleware that will Unmarshal json on data.payload and save it on data.doc
func parse(next func(*types.DispatchData) error) func(*types.DispatchData) error {
	return func(data *types.DispatchData) error {
		holder := make(map[string]interface{})
		err := json.Unmarshal([]byte(data.Payload), &holder)
		if err != nil {
			return err
		}
		data.Doc = holder
		return next(data)
	}
}
