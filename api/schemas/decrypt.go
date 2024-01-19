package schemas

type DecryptRequest struct {
	CiphertextBlob      []byte             `json:"CiphertextBlob"`
	DryRun              *bool              `json:"DryRun"`
	EncryptionAlgorithm *string            `json:"EncryptionAlgorithm"`
	EncryptionContext   *map[string]string `json:"EncryptionContext"`
	GrantTokens         *[]string          `json:"GrantTokens"`
	KeyId               *string            `json:"KeyId"`
	Recipient           *RecipientInfo     `json:"Recipient"`
}

type DecryptResponse struct {
	CiphertextForRecipient []byte `json:"CiphertextForRecipient"`
	EncryptionAlgorithm    string `json:"EncryptionAlgorithm"`
	KeyId                  string `json:"KeyId"`
	Plaintext              []byte `json:"Plaintext"`
}
