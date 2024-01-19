package schemas

type VerifyMacRequest struct {
	KeyId        string    `json:"KeyId"`
	Mac          []byte    `json:"Mac"`
	MacAlgorithm string    `json:"MacAlgorithm"`
	Message      []byte    `json:"Message"`
	DryRun       *bool     `json:"DryRun"`
	GrantTokens  *[]string `json:"GrantTokens"`
}

type VerifyMacResponse struct {
	KeyId        string `json:"KeyId"`
	MacAlgorithm string `json:"MacAlgorithm"`
	MacValid     bool   `json:"MacValid"`
}
