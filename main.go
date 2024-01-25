// server go with mux
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// create new mux
	mux := http.NewServeMux()

	// create new server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// handle route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// run server
	log.Fatal(srv.ListenAndServe())
}
