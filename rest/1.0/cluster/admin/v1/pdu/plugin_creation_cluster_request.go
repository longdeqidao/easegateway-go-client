package pdu

type PluginCreationClusterRequest struct {
	ClusterOperationRequest
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
