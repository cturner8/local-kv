package schemas

type GenerateRandomRequest struct {
	CustomKeyStoreId *string        `json:"CustomKeyStoreId"`
	NumberOfBytes    *int           `json:"NumberOfBytes"`
	Recipient        *RecipientInfo `json:"Recipient"`
}

type GenerateRandomResponse struct {
	CiphertextForRecipient []byte `json:"CiphertextForRecipient"`
	Plaintext              []byte `json:"Plaintext"`
}
