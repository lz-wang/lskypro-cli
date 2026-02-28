package config

// Config 应用配置结构
type Config struct {
	ServerURL string `yaml:"server_url" json:"server_url"`
	Token     string `yaml:"token" json:"token"`
	Output    string `yaml:"output" json:"output"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Output: "table",
	}
}
