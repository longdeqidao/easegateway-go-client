package config

type SimpleCommonCachePluginConfig struct {
	commonPluginConfig
	HitKeys     []string `json:"hit_keys"`
	CacheKey    string   `json:"cache_key"`
	TTLSec      uint32   `json:"ttl_sec"` // up to 4294967295, zero means infinite time to live
	FinishIfHit bool     `json:"finish_if_hit"`
}
