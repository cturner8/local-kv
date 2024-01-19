package schemas

type RecipientInfo struct {
	AttestationDocument    []byte `json:"AttestationDocument"`
	KeyEncryptionAlgorithm string `json:"KeyEncryptionAlgorithm"`
}
