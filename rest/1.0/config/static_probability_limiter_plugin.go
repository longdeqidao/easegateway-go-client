package config

type StaticProbabilityLimiterPluginConfig struct {
	commonPluginConfig
	PassPr float32 `json:"pass_pr"`
}
