package config

type ServiceCircuitBreakerPluginConfig struct {
	commonPluginConfig
	PluginsConcerned []string `json:"plugins_concerned"`
	// condition to enable circuit breaker
	AllTPSThresholdToEnablement float64 `json:"all_tps_threshold_to_enable"`
	// conditions to turns circuit breaker open (fully close request flow)
	FailureTPSThresholdToBreak        float64 `json:"failure_tps_threshold_to_break"`
	FailureTPSPercentThresholdToBreak float32 `json:"failure_tps_percent_threshold_to_break"`
	// condition to turns circuit breaker half-open (try service availability)
	RecoveryTimeMSec uint32 `json:"recovery_time_msec"` // up to 4294967295, equals to MTTR generally
	// condition to turns circuit breaker closed (fully open request flow)
	SuccessTPSThresholdToOpen float64 `json:"success_tps_threshold_to_open"`
}
