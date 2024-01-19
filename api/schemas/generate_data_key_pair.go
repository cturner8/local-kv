package schemas

type GenerateDataKeyPairRequest struct {
	KeyId             string             `json:"KeyId"`
	KeyPairSpec       string             `json:"KeyPairSpec"`
	DryRun            *bool              `json:"DryRun"`
	EncryptionContext *map[string]string `json:"EncryptionContext"`
	GrantTokens       *[]string          `json:"GrantTokens"`
	Recipient         *RecipientInfo     `json:"Recipient"`
}

type GenerateDataKeyPairResponse struct {
	CiphertextForRecipient   []byte `json:"CiphertextForRecipient"`
	KeyId                    string `json:"KeyId"`
	KeyPairSpec              string `json:"KeyPairSpec"`
	PrivateKeyCiphertextBlob []byte `json:"PrivateKeyCiphertextBlob"`
	PrivateKeyPlaintext      []byte `json:"PrivateKeyPlaintext"`
	PublicKey                []byte `json:"PublicKey"`
}
