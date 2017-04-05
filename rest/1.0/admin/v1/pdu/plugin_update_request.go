package pdu

type PluginUpdateRequest struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
