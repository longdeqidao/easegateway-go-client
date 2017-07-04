package pdu

import (
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type StatisticsClusterRequest struct {
	common_pdu.ClusterRequest
	Details bool `json:"details,omitempty"`
}
