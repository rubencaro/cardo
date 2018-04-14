package types

import (
	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/websocket"
)

// DispatchData is the data from th dispatcher
type DispatchData struct {
	Conn    *websocket.Conn
	Coll    driver.Collection
	Prefix  string
	Payload string
	Doc     interface{}
}
