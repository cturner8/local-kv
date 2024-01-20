package operations

import (
	"encoding/json"
	"net/http"

	"cturner8/local-kv/schemas"
)

func ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	var response = schemas.ListKeysResponse{
		Keys:      []schemas.KeyListEntry{},
		Truncated: false,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Write(jsonData)
}
