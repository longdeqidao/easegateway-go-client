package config

type NoMoreFailureLimiterPluginConfig struct {
	commonPluginConfig
	FailureCountThreshold uint64 `json:"failure_count_threshold"` // up to 18446744073709551615
	FailureTaskDataKey    string `json:"failure_task_data_key"`
	FailureTaskDataValue  string `json:"failure_task_data_value"`
}
