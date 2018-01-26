// Cluster infomation messages
package pdu

import (
	"time"
)

type ClusterResp struct {
	Time time.Time `json:"time"`
}

//
// Cluster health messages
//

type OpLogInfo struct {
	OpLogInnerInfo
	OpLogStatus
}

type OpLogGroupInfo struct {
	OpLogStatus
	UnSyncedMembersCount uint64   `json:"unsynced_members_count"`
	UnSyncedMembers      []string `json:"unsynced_members"`
}

type OpLogStatus struct {
	Synced      bool  `json:"synced"`
	MaxSequence int64 `json:"max_sequence"`
	MinSequence int64 `json:"min_sequence"`
}

type OpLogInnerInfo struct {
	DataPath     string `json:"data_path"`
	SizeBytes    uint64 `json:"size_bytes"`
	SyncProgress int8   `json:"sync_progress"`
	SyncLag      int64  `json:"sync_lag"`
}

type MemberInfo struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Mode     string `json:"mode"`
}

type MembersInfo struct {
	FailedMembersCount uint64       `json:"failed_members_count"`
	FailedMembers      []MemberInfo `json:"failed_members"`
	AliveMembersCount  uint64       `json:"alive_members_count"`
	AliveMembers       []MemberInfo `json:"alive_members"`
}

type GossipConfig struct {
	GossipPacketSizeBytes  int           `json:"gossip_packet_size_bytes"`
	GossipIntervalMillisec time.Duration `json:"gossip_interval_ms"`
}

type OpLogConfig struct {
	MaxSeqGapToPull     uint16        `json:"max_seq_gap_to_pull"`
	PullMaxCountOnce    uint16        `json:"pull_max_count_once"`
	PullIntervalSeconds time.Duration `json:"pull_interval_seconds"`
	PullTimeoutSeconds  time.Duration `json:"pull_timeout_seconds"`
}

type memberConfig struct {
	ClusterDefaultOpTimeoutSec time.Duration `json:"cluster_default_op_timeout_sec"`
	GossipConfig               GossipConfig  `json:"gossip"`
	OpLogConfig                OpLogConfig   `json:"oplog"`
}

type MemberInnerInfo struct {
	Config    memberConfig `json:"config"`
	OpLogInfo OpLogInfo    `json:"oplog"`
}
