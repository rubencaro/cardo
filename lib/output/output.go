package output

import "github.com/gorilla/websocket"

// Send simply writes given msg to given conn
func Send(conn *websocket.Conn, msg string) error {
	return conn.WriteMessage(websocket.TextMessage, []byte(msg))
}
