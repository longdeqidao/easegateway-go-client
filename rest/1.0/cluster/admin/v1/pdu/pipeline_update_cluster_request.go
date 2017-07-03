package pdu

type PipelineUpdateClusterRequest struct {
	ClusterOperationRequest
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
