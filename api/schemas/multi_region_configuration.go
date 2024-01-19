package schemas

type MultiRegionConfiguration struct {
	MultiRegionKeyType *string          `json:"MultiRegionKeyType"`
	PrimaryKey         *MultiRegionKey  `json:"PrimaryKey"`
	ReplicaKeys        []MultiRegionKey `json:"ReplicaKeys"`
}
