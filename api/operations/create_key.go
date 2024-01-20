package operations

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cturner8/local-kv/config"
	"cturner8/local-kv/constants"
	"cturner8/local-kv/crypto"
	"cturner8/local-kv/schemas"

	"github.com/google/uuid"
)

type CreateKeyController struct {
	db        *sql.DB
	masterKey []byte
}

func NewCreateKeyController(db *sql.DB, masterKey []byte) *CreateKeyController {
	return &CreateKeyController{db: db, masterKey: masterKey}
}

func (c *CreateKeyController) CreateKeyHandler(w http.ResponseWriter, r *http.Request) {
	var (
		db        *sql.DB   = c.db
		masterKey []byte    = c.masterKey
		now       time.Time = time.Now()
	)

	// key := schemas.KeyMetadata{}

	emptyString := ""
	body := schemas.KeyMetadata{
		KeySpec:               &constants.DEFAULT_KEY_SPEC,
		KeyUsage:              &constants.DEFAULT_KEY_USAGE,
		MultiRegion:           false,
		Origin:                &constants.DEFAULT_ORIGIN,
		Description:           &emptyString,
		CustomerMasterKeySpec: &emptyString,
		CustomKeyStoreId:      &emptyString,
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	// create the key metadata
	id := uuid.NewString()
	awsAccountId := "000000000000"
	arn := fmt.Sprintf("arn:aws:kms:%s:%s:key/%s", config.LOCAL_KV_REGION, awsAccountId, id)

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
		id,
		arn,
		awsAccountId, // TODO
		now,
		now,
		body.CustomerMasterKeySpec,
		body.CustomKeyStoreId,
		body.Description,
		true,
		"CUSTOMER",
		body.MultiRegion,
		body.KeySpec,
		body.KeyUsage,
		body.Origin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	// generate the key material
	dataKey := crypto.GenerateDataKey()
	additionalData := []byte(id)

	protectedDataKey := crypto.Encrypt(masterKey, dataKey, &additionalData)
	protectedKeyMaterial := crypto.Encrypt(dataKey, crypto.GenerateDataKey(), &additionalData)
	unprotectedKeyMaterial := base64.StdEncoding.Strict().EncodeToString(crypto.GenerateDataKey())

	_, err = db.Exec(
		`INSERT INTO 
			KeyMaterial (
				id,
				keyId,
				protectedDataKey,
				protectedKeyMaterial,
				unprotectedKeyMaterial
			) 
			VALUES (?, ?, ?, ?, ?)
		`,
		uuid.NewString(),
		id,
		protectedDataKey,
		protectedKeyMaterial,
		unprotectedKeyMaterial)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
