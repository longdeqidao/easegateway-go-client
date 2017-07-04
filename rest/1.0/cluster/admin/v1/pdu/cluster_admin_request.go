package pdu

import (
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type ClusterRetrieveRequest struct {
	common_pdu.ClusterRequest
	Consistent bool `json:"consistent,omitempty"`
}

type ClusterOperationRequest struct {
	ClusterRetrieveRequest
	OperationSeq uint64 `json:"operation_seq"`
}
