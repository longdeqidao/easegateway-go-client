package v1

import (
	"net/http"
	"time"
)

type Configuration struct {
	APIKeyPrefix  map[string]string `json:"APIKeyPrefix,omitempty"`
	APIKey        map[string]string `json:"APIKey,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Timeout       *time.Duration    `json:"timeout,omitempty"`
	APIClient     *APIClient
	Transport     *http.Transport
}

func NewConfiguration() *Configuration {
	timeout := 120 * time.Second

	cfg := &Configuration{
		BasePath:      "http://localhost:9090/admin/v1",
		DefaultHeader: make(map[string]string),
		APIKey:        make(map[string]string),
		APIKeyPrefix:  make(map[string]string),
		UserAgent:     "EaseGateway/go-client/1.0",
		APIClient:     &APIClient{},
		Timeout:       &timeout,
	}

	cfg.APIClient.config = cfg
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) GetAPIKeyWithPrefix(APIKeyIdentifier string) string {
	if c.APIKeyPrefix[APIKeyIdentifier] != "" {
		return c.APIKeyPrefix[APIKeyIdentifier] + " " + c.APIKey[APIKeyIdentifier]
	}

	return c.APIKey[APIKeyIdentifier]
}
