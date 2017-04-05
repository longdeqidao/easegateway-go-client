package config

type HttpHeaderCounterPluginConfig struct {
	commonPluginConfig
	HeaderConcerned string `json:"header_concerned"`
	ExpirationSec   uint32 `json:"expiration_sec"`
}
