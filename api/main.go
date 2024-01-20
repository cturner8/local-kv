package main

import (
	"database/sql"
	"log"
	"net/http"

	"cturner8/local-kv/config"
	"cturner8/local-kv/crypto"
	"cturner8/local-kv/operations"
	"cturner8/local-kv/utils"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func connectDatabase() *sql.DB {
	// Initialise database
	log.Println("Connecting to database...")

	db, err := sql.Open("sqlite3", config.LOCAL_KV_DATA_DIR+"/vault.db")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	return db
}

func setupRouter(db *sql.DB, masterKey []byte) *mux.Router {
	router := mux.NewRouter()
	router.Use(HeadersMiddleware)
	router.Use(LoggingMiddleware)
	router.Use(ErrorMiddleware)

	// Create controllers
	listKeysController := operations.NewListKeysController(db)
	createKeysController := operations.NewCreateKeyController(db, masterKey)

	// API operations
	router.HandleFunc("/", listKeysController.ListKeysHandler).Methods("POST").Headers("X-Amz-Target", "TrentService.ListKeys")
	router.HandleFunc("/", createKeysController.CreateKeyHandler).Methods("POST").Headers("X-Amz-Target", "TrentService.CreateKey")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running..."))
	})

	return router
}

func getMasterKey() []byte {
	// get the raw master key from the secret file
	rawMasterKey := utils.ReadSecretFile(config.LOCAL_KV_MASTER_KEY_FILE)
	// derive the master key
	if config.LOCAL_KV_TEMP_SALT == "" {
		panic("Error, key salt not found")
	}
	masterKey := crypto.DeriveKey(rawMasterKey, []byte(config.LOCAL_KV_TEMP_SALT)) // TODO: generate a secure salt
	return masterKey
}

func main() {
	log.Println("Initializing API...")

	config.ConfigureEnvironment()
	db := connectDatabase()
	masterKey := getMasterKey()

	router := setupRouter(db, masterKey)

	log.Println("API is running...")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}
