package schemas

type GetPublicKeyRequest struct {
	KeyId       string    `json:"KeyId"`
	GrantTokens *[]string `json:"GrantTokens"`
}

type GetPublicKeyResponse struct {
	CustomerMasterKeySpec string   `json:"CustomerMasterKeySpec"`
	EncryptionAlgorithms  []string `json:"EncryptionAlgorithms"`
	KeyId                 string   `json:"KeyId"`
	KeySpec               string   `json:"KeySpec"`
	KeyUsage              string   `json:"KeyUsage"`
	PublicKey             []byte   `json:"PublicKey"`
	SigningAlgorithms     []string `json:"SigningAlgorithms"`
}
