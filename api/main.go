package main

import (
	"database/sql"
	"log"
	"net/http"

	"cturner8/local-kv/operations"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Initializing API...")

	router := mux.NewRouter()
	router.Use(HeadersMiddleware)
	router.Use(LoggingMiddleware)

	// Initialise database
	log.Println("Connecting to database...")

	db, err := sql.Open("sqlite3", "./vault.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	// API operations
	router.HandleFunc("/", operations.ListKeysHandler).Methods("POST").Headers("X-Amz-Target", "TrentService.ListKeys")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running..."))
	})

	log.Println("API is running...")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
