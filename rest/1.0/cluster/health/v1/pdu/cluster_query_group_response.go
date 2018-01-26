package pdu

type ClusterQueryGroupResponse struct {
	ClusterResp
	OpLogGroupInfo `json:"oplog"`
	MembersInfo
}

type ClusterQueryGroupsResponse struct {
	ClusterResp
	GroupsCount int      `json:"groups_count"`
	Groups      []string `json:"groups"`
}
