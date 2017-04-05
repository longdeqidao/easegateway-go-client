package config

type EaseMonitorGraphiteValidatorPluginConfig struct {
	commonPluginConfig
	DataKey string `json:"data_key"`
}
