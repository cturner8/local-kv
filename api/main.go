package main

import (
	"cturner8/local-kv/operations"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Initializing API")

	router := mux.NewRouter()
	router.Use(HeadersMiddleware)
	router.Use(LoggingMiddleware)

	// API operations
	router.HandleFunc("/", operations.ListKeysHandler).Methods("POST").Headers("X-Amz-Target", "TrentService.ListKeys")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running..."))
	})

	log.Println("API is running...")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
