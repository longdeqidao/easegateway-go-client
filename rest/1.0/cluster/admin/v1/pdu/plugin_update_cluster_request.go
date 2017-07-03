package pdu

type PluginUpdateClusterRequest struct {
	ClusterOperationRequest
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
