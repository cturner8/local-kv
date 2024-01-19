package schemas

type GenerateDataKeyPairWithoutPlaintextRequest struct {
	KeyId             string             `json:"KeyId"`
	KeyPairSpec       string             `json:"KeyPairSpec"`
	DryRun            *bool              `json:"DryRun"`
	EncryptionContext *map[string]string `json:"EncryptionContext"`
	GrantTokens       *[]string          `json:"GrantTokens"`
}

type GenerateDataKeyPairWithoutPlaintextResponse struct {
	KeyId                    string `json:"KeyId"`
	KeyPairSpec              string `json:"KeyPairSpec"`
	PrivateKeyCiphertextBlob []byte `json:"PrivateKeyCiphertextBlob"`
	PublicKey                []byte `json:"PublicKey"`
}
