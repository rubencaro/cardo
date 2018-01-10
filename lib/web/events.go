package web

import (
	"fmt"
	"log"
	"net/http"
)

// Messages is the channel where you should send your strings
// They will be sent to the client connected
var Messages = make(chan string)

// EventsHandler handles the connection and dispatching
// of events to the client
func EventsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Cannot stream", http.StatusInternalServerError)
		return
	}

	cn, ok := w.(http.CloseNotifier)
	if !ok {
		http.Error(w, "Cannot notify closure", http.StatusInternalServerError)
		return
	}

	f.Flush()

	log.Println("SSE connected")

	for {
		select {
		case <-cn.CloseNotify():
			log.Println("SSE closed connection")
			return
		case msg := <-Messages:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			fmt.Printf("SSE sent: %s", msg)
			f.Flush()
		}
	}
}
