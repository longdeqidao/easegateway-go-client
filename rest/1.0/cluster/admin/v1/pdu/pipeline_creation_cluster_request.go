package pdu

type PipelineCreationClusterRequest struct {
	ClusterOperationRequest
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
