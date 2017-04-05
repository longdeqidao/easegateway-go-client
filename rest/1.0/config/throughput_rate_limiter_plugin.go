package config

type ThroughputRateLimiterPluginConfig struct {
	commonPluginConfig
	Tps string `json:"tps,omitempty"`
}
