package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal([]string{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// Middleware function, which will be called for each request
func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Initializing API")

	router := mux.NewRouter()
	router.Use(HeadersMiddleware)

	// API operations
	router.HandleFunc("/", ListKeysHandler).Methods("POST").Headers("X-Amz-Target", "TrentService.ListKeys")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running..."))
	})

	log.Println("API is running...")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
