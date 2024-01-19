package schemas

type DescribeKeyRequest struct {
	KeyId       string    `json:"KeyId"`
	GrantTokens *[]string `json:"GrantTokens"`
}

type DescribeKeyResponse struct {
	KeyMetadata KeyMetadata `json:"KeyMetadata"`
}
