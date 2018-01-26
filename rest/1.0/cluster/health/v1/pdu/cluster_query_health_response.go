package pdu

type ClusterQueryGroupHealthResponse struct {
	ClusterResp
	Status      string `json:"status"`
	Description string `json:"description"`
}

type ClusterQueryGroupsHealthResponse struct {
	ClusterQueryGroupHealthResponse
}
