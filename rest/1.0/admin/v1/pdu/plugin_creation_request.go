package pdu

type PluginCreationRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
