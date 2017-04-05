package pdu

type PluginsRetrieveRequest struct {
	NamePattern string   `json:"name_pattern,omitempty"`
	Types       []string `json:"types,omitempty"`
}
