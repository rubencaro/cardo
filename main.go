package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rubencaro/cardo/lib/web"
)

func main() {

	http.HandleFunc("/ping", web.PingHandler)
	http.HandleFunc("/events", web.EventsHandler)
	http.HandleFunc("/ws", web.SocketsHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	// start server
	srv := &http.Server{
		Addr:         "localhost:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
