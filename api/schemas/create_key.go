package schemas

type CreateKeyRequest struct {
	BypassPolicyLockoutSafetyCheck bool    `json:"BypassPolicyLockoutSafetyCheck"`
	CustomerMasterKeySpec          *string `json:"CustomerMasterKeySpec"`
	CustomKeyStoreId               *string `json:"CustomKeyStoreId"`
	Description                    string  `json:"Description"`
	KeySpec                        string  `json:"KeySpec"`
	KeyUsage                       string  `json:"KeyUsage"`
	MultiRegion                    bool    `json:"MultiRegion"`
	Origin                         *string `json:"Origin"`
	Policy                         *string `json:"Policy"`
	Tags                           *[]Tag  `json:"Tags"`
	XksKeyId                       string  `json:"XksKeyId"`
}

type CreateKeyResponse struct {
	KeyMetadata KeyMetadata `json:"KeyMetadata"`
}
