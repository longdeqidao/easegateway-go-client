package pdu

type ClusterRequest struct {
	TimeoutSec int32 `json:"timeout_sec,omitempty"`
}

type ClusterRetrieveRequest struct {
	ClusterRequest
	Consistent bool `json:"consistent,omitempty"`
}

type ClusterOperationRequest struct {
	ClusterRetrieveRequest
	OperationSeq int32 `json:"operation_seq"`
}
