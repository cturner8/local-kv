package operations

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	// "time"

	"cturner8/local-kv/constants"
	"cturner8/local-kv/crypto/aes"
	"cturner8/local-kv/crypto/rsa"
	"cturner8/local-kv/crypto/x509"
	"cturner8/local-kv/schemas"

	"golang.org/x/exp/slices"
)

type EncryptController struct {
	db        *sql.DB
	masterKey []byte
}

var ENCRYPT_HEADER = "TrentService.Encrypt"

type KeyMetadataWithMaterial struct {
	KeyId                  string
	KeySpec                string
	KeyUsage               string
	ProtectedDataKey       string
	ProtectedKeyMaterial   string
	UnprotectedKeyMaterial *string
}

func NewEncryptController(db *sql.DB, masterKey []byte) *EncryptController {
	return &EncryptController{db: db, masterKey: masterKey}
}

func validateEncryptRequest(w http.ResponseWriter, body schemas.EncryptRequest) {
	log.Print("Validating request...")

	log.Print("Validating user provided values")

	if !slices.Contains(constants.ENCRYPTION_ALGORITHMS, *body.EncryptionAlgorithm) {
		err := errors.New("invalid algorithm provided")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	if len(body.Plaintext) == 0 {
		err := errors.New("invalid plaintext")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	log.Print("Request validated")
}

func (c *EncryptController) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var (
		db        *sql.DB = c.db
		masterKey []byte  = c.masterKey
	)

	body := schemas.EncryptRequest{
		EncryptionAlgorithm: &constants.DEFAULT_ENCRYPTION_ALGORITHM,
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	validateEncryptRequest(w, body)

	var keyMetadata KeyMetadataWithMaterial
	row := db.QueryRow(
		`SELECT
			k.id,
			k.keySpec,
			k.keyUsage,
			km.protectedDataKey,
			km.protectedKeyMaterial,
			km.unprotectedKeyMaterial
		FROM
			KeyMetadata k
		LEFT JOIN
			KeyMaterial km
			ON km.keyId = k.id
		WHERE
			k.id = ?
		`,
		body.KeyId)

	if err := row.Scan(
		&keyMetadata.KeyId,
		&keyMetadata.KeySpec,
		&keyMetadata.KeyUsage,
		&keyMetadata.ProtectedDataKey,
		&keyMetadata.ProtectedKeyMaterial,
		&keyMetadata.UnprotectedKeyMaterial); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	var ciphertext string
	encryptionContext := []byte{}

	plaintext, err := base64.StdEncoding.DecodeString(body.Plaintext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	if keyMetadata.KeyUsage != constants.ENCRYPT_DECRYPT {
		err := errors.New("key not permitted for encryption")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	if slices.Contains(constants.SYMMETRIC_KEY_SPECS, keyMetadata.KeySpec) {
		additionalData := []byte(keyMetadata.KeyId)
		dataKey := aes.Decrypt(masterKey, keyMetadata.ProtectedDataKey, &additionalData)

		keyMaterial := aes.Decrypt(dataKey, keyMetadata.ProtectedKeyMaterial, &additionalData)
		ciphertext = aes.Encrypt(keyMaterial, plaintext, &encryptionContext)
	} else {
		base64KeyMaterial, err := base64.StdEncoding.DecodeString(*keyMetadata.UnprotectedKeyMaterial)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		publicKey := x509.DecodePublicKey(base64KeyMaterial)
		ciphertext = rsa.Encrypt(publicKey, plaintext, encryptionContext)
	}

	response := schemas.EncryptResponse{
		CiphertextBlob:      ciphertext,
		EncryptionAlgorithm: *body.EncryptionAlgorithm,
		KeyId:               body.KeyId,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write(jsonData)
}
