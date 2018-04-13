package input

import (
	"fmt"
	"regexp"

	driver "github.com/arangodb/go-driver"
	"github.com/gorilla/websocket"
)

var logRE = regexp.MustCompile(`^log: `)

// Dispatch reads incoming message from given Conn and dispatches it
// to the right function
func Dispatch(conn *websocket.Conn, coll driver.Collection, msg string) error {
	switch {
	case logRE.MatchString(msg):
		handleLog(coll, msg)
		return nil
	default:
		conn.Close()
		return fmt.Errorf("Unexpected msg: %s", msg)
	}
}
