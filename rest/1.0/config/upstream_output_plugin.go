package config

type UpstreamOutputPluginConfig struct {
	commonPluginConfig
	TargetPipelineNames []string `json:"target_pipelines"`
	RoutePolicy         string   `json:"route_policy"`
	TimeoutSec          uint16   `json:"timeout_sec"` // up to 65535, zero means no timeout
	RequestDataKeys     []string `json:"request_data_keys"`
	// for weighted_round_robin and weighted_random policies
	TargetPipelineWeights []uint16 `json:"target_weights"` // weight up to 65535, 0 based
	// for hash policy
	ValueHashedKeys []string `json:"value_hashed_keys"`
	// for filter policy
	// each map in the list as the condition set for the target pipeline according to the index
	// map key is the key of value in the task, map value is the match condition, support regex
	FilterConditions []map[string]string `json:"filter_conditions"`
}
