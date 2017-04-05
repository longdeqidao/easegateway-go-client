package config

type HTTPInputPluginConfig struct {
	commonPluginConfig
	URL         string              `json:"url"`
	Method      string              `json:"method"`
	HeadersEnum map[string][]string `json:"headers_enum"`
	Unzip       bool                `json:"unzip"`
	RespondErr  bool                `json:"respond_error"`
	FastClose   bool                `json:"fast_close"`

	RequestHeaderNamesKey string `json:"request_header_names_key"`
	RequestBodyIOKey      string `json:"request_body_io_key"`
	ResponseCodeKey       string `json:"response_code_key"`
	ResponseBodyIOKey     string `json:"response_body_io_key"`
	ResponseBodyBufferKey string `json:"response_body_buffer_key"`
}
