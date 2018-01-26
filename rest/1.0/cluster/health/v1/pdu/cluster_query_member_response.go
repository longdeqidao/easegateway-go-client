package pdu

type ClusterQueryMemberResponse struct {
	ClusterResp
	MemberInnerInfo
}

type ClusterQueryMembersResponse struct {
	ClusterResp
	MembersInfo
}
