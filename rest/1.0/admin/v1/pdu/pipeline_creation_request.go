package pdu

type PipelineCreationRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
