package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Initializing API")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running..."))
	})

	log.Println("API is running...")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
