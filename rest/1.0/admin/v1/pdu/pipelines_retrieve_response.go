package pdu

type PipelinesRetrieveResponse struct {
	Pipelines []PipelineSpec `json:"pipelines"`
}
