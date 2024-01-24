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

var ROUTER_HEADER_NAME = "X-Amz-Target"

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
	encryptController := operations.NewEncryptController(db, masterKey)

	// API operations
	router.HandleFunc("/", listKeysController.ListKeysHandler).Methods("POST").Headers(ROUTER_HEADER_NAME, operations.LIST_KEYS_HEADER)
	router.HandleFunc("/", createKeysController.CreateKeyHandler).Methods("POST").Headers(ROUTER_HEADER_NAME, operations.CREATE_KEY_HEADER)
	router.HandleFunc("/", encryptController.EncryptHandler).Methods("POST").Headers(ROUTER_HEADER_NAME, operations.ENCRYPT_HEADER)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running..."))
	})

	return router
}

func getMasterKey() []byte {
	// get the raw master key and salt from the secret file
	rawMasterKey := utils.ReadSecretFile(config.LOCAL_KV_MASTER_KEY_FILE)
	salt := utils.ReadSecretFile(config.LOCAL_KV_MASTER_KEY_SALT_FILE)
	// derive the master key
	masterKey := crypto.DeriveKey(rawMasterKey, salt)
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
