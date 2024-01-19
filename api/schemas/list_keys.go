package schemas

type ListKeysRequest struct {
	Limit  int    `json:"Limit"`
	Marker string `json:"Marker"`
}

type ListKeysResponse struct {
	Keys       []KeyListEntry `json:"Keys"`
	NextMarker *string        `json:"NextMarker"`
	Truncated  bool           `json:"Truncated"`
}
