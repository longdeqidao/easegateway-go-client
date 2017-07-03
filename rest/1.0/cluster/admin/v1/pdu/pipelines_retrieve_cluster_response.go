package pdu

type PipelinesRetrieveClusterResponse struct {
	Pipelines []PipelineSpec `json:"pipelines"`
}
