package pdu

type PluginSpec struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}
