package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rubencaro/cardo/lib/cnf"
	"github.com/rubencaro/cardo/lib/db"
	"github.com/rubencaro/cardo/lib/web"
)

func main() {
	c, err := cnf.Read()
	if err != nil {
		fmt.Println("Error reading config: ", err)
		return
	}

	logs, err := db.GetCollection("logs", &c)
	if err != nil {
		fmt.Println("Error getting database: ", err)
		return
	}

	http.HandleFunc("/ping", web.PingHandler)
	http.HandleFunc("/events", web.EventsHandler)
	http.HandleFunc("/ws", web.SocketsHandler(logs))
	http.Handle("/", http.FileServer(http.Dir("static")))

	// start server
	srv := &http.Server{
		Addr:         "localhost:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening at %s...\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
