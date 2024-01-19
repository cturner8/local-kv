package schemas

type SignRequest struct {
	KeyId            string    `json:"KeyId"`
	Message          []byte    `json:"Message"`
	SigningAlgorithm string    `json:"SigningAlgorithm"`
	DryRun           *bool     `json:"DryRun"`
	GrantTokens      *[]string `json:"GrantTokens"`
	MessageType      *string   `json:"MessageType"`
}

type SignResponse struct {
	KeyId            string `json:"KeyId"`
	Signature        []byte `json:"Signature"`
	SigningAlgorithm string `json:"SigningAlgorithm"`
}
