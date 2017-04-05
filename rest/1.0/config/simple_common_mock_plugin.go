package config

type SimpleCommonMockPluginConfig struct {
	commonPluginConfig
	PluginConcerned        string `json:"plugin_concerned"`
	TaskErrorCodeConcerned string `json:"task_error_code_concerned"`
	MockTaskDataKey        string `json:"mock_task_data_key"`
	MockTaskDataValue      string `json:"mock_task_data_value"`
}
