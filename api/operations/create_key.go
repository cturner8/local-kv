package operations

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"cturner8/local-kv/config"
	"cturner8/local-kv/constants"
	"cturner8/local-kv/crypto/aes"
	"cturner8/local-kv/crypto/rsa"
	"cturner8/local-kv/crypto/x509"
	"cturner8/local-kv/schemas"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type CreateKeyController struct {
	db        *sql.DB
	masterKey []byte
}

var CREATE_KEY_HEADER = "TrentService.CreateKey"

func NewCreateKeyController(db *sql.DB, masterKey []byte) *CreateKeyController {
	return &CreateKeyController{db: db, masterKey: masterKey}
}

func validateRequest(w http.ResponseWriter, body schemas.CreateKeyRequest) {
	log.Print("Validating request...")

	log.Print("Validating user provided values")
	if !slices.Contains(constants.ORIGINS, *body.Origin) {
		err := errors.New("invalid origin provided")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	if !slices.Contains(constants.KEY_SPECS, body.KeySpec) {
		err := errors.New("invalid keyspec provided")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	if !slices.Contains(constants.KEY_USAGES, body.KeyUsage) {
		err := errors.New("invalid keyusage provided")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	log.Print("Validating key usage is appropriate for the key spec")

	// now that we've validated the key spec, we can check the key usage is appropriate for the key spec

	// symmetric keys
	if slices.Contains(constants.SYMMETRIC_KEY_SPECS, body.KeySpec) {
		if !slices.Contains(constants.SYMMETRIC_ENCRYPTION_KEY_USAGES, body.KeyUsage) {
			err := errors.New("invalid symmetric encryption keyusage provided")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	}

	// asymmetric keys
	if slices.Contains(constants.ASYMMETRIC_KEY_SPECS, body.KeySpec) {
		if !slices.Contains(constants.ASYMMETRIC_ENCRYPTION_KEY_USAGES, body.KeyUsage) {
			err := errors.New("invalid asymmetric encryption keyusage provided")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	}

	log.Print("Request validated")
}

func (c *CreateKeyController) CreateKeyHandler(w http.ResponseWriter, r *http.Request) {
	var (
		db        *sql.DB   = c.db
		masterKey []byte    = c.masterKey
		now       time.Time = time.Now()
	)

	// key := schemas.KeyMetadata{}

	emptyString := ""
	body := schemas.CreateKeyRequest{
		KeySpec:               constants.DEFAULT_KEY_SPEC,
		KeyUsage:              constants.DEFAULT_KEY_USAGE,
		MultiRegion:           false,
		Origin:                &constants.DEFAULT_ORIGIN,
		Description:           emptyString,
		CustomerMasterKeySpec: &emptyString,
		CustomKeyStoreId:      &emptyString,
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	validateRequest(w, body)

	// create the key metadata
	id := uuid.NewString()
	awsAccountId := "000000000000" // TODO
	arn := fmt.Sprintf("arn:aws:kms:%s:%s:key/%s", config.LOCAL_KV_REGION, awsAccountId, id)

	keyMetadata := schemas.KeyMetadata{
		KeyId:                 id,
		Arn:                   &arn,
		AWSAccountId:          &awsAccountId,
		Description:           &body.Description,
		Enabled:               true,
		KeyManager:            "CUSTOMER",
		MultiRegion:           body.MultiRegion,
		KeySpec:               &body.KeySpec,
		KeyUsage:              &body.KeyUsage,
		Origin:                body.Origin,
		CustomerMasterKeySpec: body.CustomerMasterKeySpec,
		CustomKeyStoreId:      body.CustomKeyStoreId,
	}

	_, err = db.Exec(
		`INSERT INTO 
			KeyMetadata (
				id,
				arn,
				awsAccountId,
				createdDate,
				updatedDate,
				customerMasterKeySpec,
				customKeyStoreId,
				description,
				enabled,
				keyManager,
				multiRegion,
				keySpec,
				keyUsage,
				origin
			) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		keyMetadata.KeyId,
		keyMetadata.Arn,
		keyMetadata.AWSAccountId,
		now,
		now,
		keyMetadata.CustomerMasterKeySpec,
		keyMetadata.CustomKeyStoreId,
		keyMetadata.Description,
		keyMetadata.Enabled,
		keyMetadata.KeyManager,
		keyMetadata.MultiRegion,
		keyMetadata.KeySpec,
		keyMetadata.KeyUsage,
		keyMetadata.Origin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	log.Printf("Created key [%s]", id)

	// generate the key material
	dataKey := aes.GenerateDataKey()
	additionalData := []byte(id)

	protectedDataKey := aes.Encrypt(masterKey, dataKey, &additionalData)
	var (
		protectedKeyMaterial   string
		unprotectedKeyMaterial *string
	)

	if slices.Contains(constants.SYMMETRIC_KEY_SPECS, body.KeySpec) {
		protectedKeyMaterial = aes.Encrypt(dataKey, aes.GenerateDataKey(), &additionalData)
	} else {
		var bits int
		if body.KeySpec == constants.RSA_2048 {
			bits = 2048
		} else if body.KeySpec == constants.RSA_3072 {
			bits = 3072
		} else if body.KeySpec == constants.RSA_4096 {
			bits = 4096
		}

		// generate a new keypair using the provide key spec bits
		privateKey, publicKey := rsa.GenerateKeyPair(bits)

		// encode the keypair for storage
		encodedPrivateKey := x509.EncodePrivateKey(privateKey)
		encodedPublicKey := x509.EncodePublicKey(publicKey)

		// encrypt the private key
		protectedKeyMaterial = aes.Encrypt(dataKey, encodedPrivateKey, &additionalData)

		// no need to encrypt the public key, that can just be base64 encoded
		unprotectedKeyMaterial = new(string)
		*unprotectedKeyMaterial = base64.StdEncoding.Strict().EncodeToString(encodedPublicKey)
	}

	_, err = db.Exec(
		`INSERT INTO 
			KeyMaterial (
				id,
				createdDate,
				updatedDate,
				keyId,
				protectedDataKey,
				protectedKeyMaterial,
				unprotectedKeyMaterial
			) 
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`,
		uuid.NewString(),
		now,
		now,
		id,
		protectedDataKey,
		protectedKeyMaterial,
		unprotectedKeyMaterial)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	log.Printf("Created [%s] key material for [%s]", body.KeySpec, id)

	response := schemas.CreateKeyResponse{
		KeyMetadata: keyMetadata,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write(jsonData)
}
