package pdu

type PipelineUpdateRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
