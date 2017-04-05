package config

type KafkaOutputPluginConfig struct {
	commonPluginConfig
	Brokers  []string `json:"brokers"`
	ClientID string   `json:"client_id"`
	Topic    string   `json:"topic"`

	MessageKeyKey string `json:"message_key_key"`
	DataKey       string `json:"data_key"`
}
