package schemas

type GenerateDataKeyWithoutPlaintextRequest struct {
	KeyId             string             `json:"KeyId"`
	DryRun            *bool              `json:"DryRun"`
	EncryptionContext *map[string]string `json:"EncryptionContext"`
	GrantTokens       *[]string          `json:"GrantTokens"`
	KeySpec           *string            `json:"KeySpec"`
	NumberOfBytes     *int               `json:"NumberOfBytes"`
}

type GenerateDataKeyWithoutPlaintextResponse struct {
	CiphertextBlob []byte `json:"CiphertextBlob"`
	KeyId          string `json:"KeyId"`
}
