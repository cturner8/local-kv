package schemas

type KeyMetadata struct {
	KeyId                       string                    `json:"KeyId"`
	Arn                         *string                   `json:"Arn"`
	AWSAccountId                *string                   `json:"AWSAccountId"`
	CloudHsmClusterId           *string                   `json:"CloudHsmClusterId"`
	CreationDate                *int                      `json:"CreationDate"`
	CustomerMasterKeySpec       *string                   `json:"CustomerMasterKeySpec"`
	CustomKeyStoreId            *string                   `json:"CustomKeyStoreId"`
	DeletionDate                *int                      `json:"DeletionDate"`
	Description                 *string                   `json:"Description"`
	Enabled                     bool                      `json:"Enabled"`
	EncryptionAlgorithms        []string                  `json:"EncryptionAlgorithms"`
	ExpirationModel             *string                   `json:"ExpirationModel"`
	KeyManager                  string                    `json:"KeyManager"`
	KeySpec                     *string                   `json:"KeySpec"`
	KeyState                    *string                   `json:"KeyState"`
	KeyUsage                    *string                   `json:"KeyUsage"`
	MacAlgorithms               *[]string                 `json:"MacAlgorithms"`
	MultiRegion                 bool                      `json:"MultiRegion"`
	MultiRegionConfiguration    *MultiRegionConfiguration `json:"MultiRegionConfiguration"`
	Origin                      *string                   `json:"Origin"`
	PendingDeletionWindowInDays *int                      `json:"PendingDeletionWindowInDays"`
	SigningAlgorithms           []string                  `json:"SigningAlgorithms"`
	ValidTo                     *int                      `json:"ValidTo"`
	XksKeyConfiguration         *XksKeyConfigurationType  `json:"XksKeyConfiguration"`
}
