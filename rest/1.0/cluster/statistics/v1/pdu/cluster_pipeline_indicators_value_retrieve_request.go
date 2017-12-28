package pdu

type ClusterPipelineIndicatorsValueRequest struct {
	StatisticsClusterRequest
	IndicatorNames []string `json:"indicator_names"`
}
