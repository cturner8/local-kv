package operations

import (
	"cturner8/local-kv/schemas"
	"encoding/json"
	"net/http"
)

func ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	var response = schemas.ListKeysResponse{
		Keys:      []schemas.KeyListEntry{},
		Truncated: false,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
