package schemas

type EncryptRequest struct {
	KeyId               string             `json:"KeyId"`
	Plaintext           []byte             `json:"Plaintext"`
	DryRun              *bool              `json:"DryRun"`
	EncryptionAlgorithm *string            `json:"EncryptionAlgorithm"`
	EncryptionContext   *map[string]string `json:"EncryptionContext"`
	GrantTokens         *[]string          `json:"GrantTokens"`
}

type EncryptResponse struct {
	CiphertextBlob      []byte `json:"CiphertextBlob"`
	EncryptionAlgorithm string `json:"EncryptionAlgorithm"`
	KeyId               string `json:"KeyId"`
}
