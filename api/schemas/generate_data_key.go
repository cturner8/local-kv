package schemas

type GenerateDataKeyRequest struct {
	KeyId             string             `json:"KeyId"`
	DryRun            *bool              `json:"DryRun"`
	EncryptionContext *map[string]string `json:"EncryptionContext"`
	GrantTokens       *[]string          `json:"GrantTokens"`
	KeySpec           *string            `json:"KeySpec"`
	NumberOfBytes     *int               `json:"NumberOfBytes"`
	Recipient         *RecipientInfo     `json:"Recipient"`
}

type GenerateDataKeyResponse struct {
	CiphertextBlob         []byte `json:"CiphertextBlob"`
	CiphertextForRecipient []byte `json:"CiphertextForRecipient"`
	KeyId                  string `json:"KeyId"`
	Plaintext              []byte `json:"Plaintext"`
}
