package schemas

type GenerateMacRequest struct {
	KeyId        string    `json:"KeyId"`
	MacAlgorithm string    `json:"MacAlgorithm"`
	Message      []byte    `json:"Message"`
	DryRun       *bool     `json:"DryRun"`
	GrantTokens  *[]string `json:"GrantTokens"`
}

type GenerateMacResponse struct {
	KeyId        string `json:"KeyId"`
	Mac          []byte `json:"Mac"`
	MacAlgorithm string `json:"MacAlgorithm"`
}
