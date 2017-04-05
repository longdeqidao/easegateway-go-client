package config

type JSONValidatorPluginConfig struct {
	commonPluginConfig
	Schema  string `json:"schema"`
	DataKey string `json:"data_key"`
}
