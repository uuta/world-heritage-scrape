package main

import (
	"log"
	"net/http"
	"time"
	"world-heritage-scrape/app/worldheritage/interfaces/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/world-heritage", handler.PutWorldHeritageHandler).Methods("PUT")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
