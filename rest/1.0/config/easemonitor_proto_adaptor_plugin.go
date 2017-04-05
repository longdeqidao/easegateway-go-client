package config

type EaseMonitorProtoAdaptorPluginConfig struct {
	commonPluginConfig
	GidKey  string `json:"gid_key"`
	DataKey string `json:"data_key"`
}
