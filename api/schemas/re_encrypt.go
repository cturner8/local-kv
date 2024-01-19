package schemas

type ReEncryptRequest struct {
	CiphertextBlob                 []byte             `json:"CiphertextBlob"`
	DestinationKeyId               string             `json:"DestinationKeyId"`
	DestinationEncryptionAlgorithm *string            `json:"DestinationEncryptionAlgorithm"`
	DestinationEncryptionContext   *map[string]string `json:"DestinationEncryptionContext"`
	DryRun                         *bool              `json:"DryRun"`
	GrantTokens                    *[]string          `json:"GrantTokens"`
	SourceEncryptionAlgorithm      *string            `json:"SourceEncryptionAlgorithm"`
	SourceEncryptionContext        *map[string]string `json:"SourceEncryptionContext"`
	SourceKeyId                    *string            `json:"SourceKeyId"`
}

type ReEncryptResponse struct {
	CiphertextBlob                 []byte `json:"CiphertextBlob"`
	DestinationEncryptionAlgorithm string `json:"DestinationEncryptionAlgorithm"`
	KeyId                          string `json:"KeyId"`
	SourceEncryptionAlgorithm      string `json:"SourceEncryptionAlgorithm"`
	SourceKeyId                    string `json:"SourceKeyId"`
}
