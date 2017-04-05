package config

type IOReaderPluginConfig struct {
	commonPluginConfig
	LengthMax int64 `json:"read_length_max"` // up to 9223372036854775807 ~= 8192 Pebibyte
	Close     bool  `json:"close_after_read"`

	InputKey  string `json:"input_key"`
	OutputKey string `json:"output_key"`
}
