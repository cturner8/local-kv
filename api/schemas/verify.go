package schemas

type VerifyRequest struct {
	KeyId            string    `json:"KeyId"`
	Message          []byte    `json:"Message"`
	Signature        []byte    `json:"Signature"`
	SigningAlgorithm string    `json:"SigningAlgorithm"`
	DryRun           *bool     `json:"DryRun"`
	GrantTokens      *[]string `json:"GrantTokens"`
	MessageType      *string   `json:"MessageType"`
}

type VerifyResponse struct {
	KeyId            string `json:"KeyId"`
	SignatureValid   bool   `json:"SignatureValid"`
	SigningAlgorithm string `json:"SigningAlgorithm"`
}
