package pdu

type PipelinesRetrieveClusterRequest struct {
	ClusterRetrieveRequest
	NamePattern string   `json:"name_pattern,omitempty"`
	Types       []string `json:"types,omitempty"`
}
