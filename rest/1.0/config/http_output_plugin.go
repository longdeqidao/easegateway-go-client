package config

type HTTPOutputPluginConfig struct {
	commonPluginConfig
	URLPattern               string            `json:"url_pattern"`
	HeaderPatterns           map[string]string `json:"header_patterns"`
	RequestBodyIOKey         string            `json:"request_body_io_key"`
	RequestBodyBufferPattern string            `json:"request_body_buffer_pattern"`
	Method                   string            `json:"method"`
	TimeoutSec               uint16            `json:"timeout_sec"` // up to 65535, zero means no timeout
	CertFile                 string            `json:"cert_file"`
	KeyFile                  string            `json:"key_file"`
	CAFile                   string            `json:"ca_file"`
	Insecure                 bool              `json:"insecure_tls"`

	ResponseCodeKey   string `json:"response_code_key"`
	ResponseBodyIOKey string `json:"response_body_io_key"`
}
