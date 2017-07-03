package pdu

type PipelineSpec struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
