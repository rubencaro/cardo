package web

import "net/http"

// PingHandler handles the ping request
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
