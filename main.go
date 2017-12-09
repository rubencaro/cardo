package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rubencaro/cardo/lib/web"
)

func main() {

	http.HandleFunc("/ping", web.PingHandler)
	http.HandleFunc("/", http.NotFound)

	// start server
	srv := &http.Server{
		Addr:         "localhost:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
