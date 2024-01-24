package operations

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"cturner8/local-kv/schemas"
)

type ListKeysController struct {
	db *sql.DB
}

var LIST_KEYS_HEADER = "TrentService.ListKeys"

func NewListKeysController(db *sql.DB) *ListKeysController {
	return &ListKeysController{db: db}
}

func (c *ListKeysController) ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	keys := []schemas.KeyListEntry{}

	rows, err := c.db.Query("SELECT id, arn FROM KeyMetadata")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		key := schemas.KeyListEntry{
			KeyId:  "",
			KeyArn: "",
		}
		if err := rows.Scan(&key.KeyId, &key.KeyArn); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		keys = append(keys, key)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	response := schemas.ListKeysResponse{
		Keys:      keys,
		Truncated: false,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write(jsonData)
}
