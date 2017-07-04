package pdu

type ClusterOperationSeqRetrieveResponse struct {
	ClusterGroup string `json:"cluster_group,omitempty"`
	OperationSeq uint64 `json:"operation_seq,omitempty"`
}
