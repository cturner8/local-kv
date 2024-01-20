package config

import (
	"encoding/base64"
	"log"
	"os"

	"cturner8/local-kv/crypto"
)

var (
	// general config
	LOCAL_KV_REGION = getEnvWithDefault("LOCAL_KV_REGION", "us-east-1")
	// data config
	LOCAL_KV_DATA_DIR = getEnvWithDefaultDirectory("LOCAL_KV_DATA_DIR", "/.local-kv/data")
	// secrets config
	LOCAL_KV_SECRETS_ENGINE  = getEnvWithDefault("LOCAL_KV_SECRETS_ENGINE", "file")
	LOCAL_KV_SECRETS_DIR     = getEnvWithDefaultDirectory("LOCAL_KV_SECRETS_DIR", "/.local-kv/secrets")
	LOCAL_KV_MASTER_KEY_FILE = LOCAL_KV_SECRETS_DIR + "/master.key"
	LOCAL_KV_TEMP_SALT       = os.Getenv("LOCAL_KV_TEMP_SALT")
)

func getEnvWithDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvWithDefaultDirectory(key, defaultDirectory string) string {
	userHome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return getEnvWithDefault(key, userHome+defaultDirectory)
}

func ConfigureEnvironment() {
	log.Print("Configuring API...")

	log.Print("Initializing storage directories...")
	err := os.MkdirAll(LOCAL_KV_DATA_DIR, 0o750)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(LOCAL_KV_SECRETS_DIR, 0o750)
	if err != nil {
		panic(err)
	}

	log.Print("Initializing secrets engine...")
	if LOCAL_KV_SECRETS_ENGINE == "file" {
		log.Print("Initializing file secrets engine...")

		_, err := os.Stat(LOCAL_KV_MASTER_KEY_FILE)
		if err != nil {
			log.Print("Generating new master key")

			keyFile, err := os.Create(LOCAL_KV_MASTER_KEY_FILE)
			if err != nil {
				panic(err)
			}
			defer keyFile.Close()

			masterKey := crypto.GenerateDataKey()
			if _, err := keyFile.WriteString(base64.StdEncoding.EncodeToString(masterKey)); err != nil {
				panic(err)
			}

			log.Print("Master key generated")
		}
	}

	log.Print("API configured")
}
