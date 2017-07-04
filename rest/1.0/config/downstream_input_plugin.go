package config

type DownstreamInputPluginConfig struct {
	commonPluginConfig
	ResponseDataKeys []string `json:"response_data_keys"`
}
