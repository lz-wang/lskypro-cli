package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// DefaultConfigDir 默认配置目录
const DefaultConfigDir = ".lc"

// DefaultConfigFile 默认配置文件名
const DefaultConfigFile = "config.yaml"

// Provider 配置提供者
type Provider struct {
	configPath string
	config     *Config
}

// NewProvider 创建配置提供者
func NewProvider(configPath string) (*Provider, error) {
	if configPath == "" {
		configPath = GetDefaultConfigPath()
	}

	p := &Provider{
		configPath: configPath,
		config:     DefaultConfig(),
	}

	if err := p.Load(); err != nil {
		return nil, err
	}

	return p, nil
}

// Load 加载配置文件
func (p *Provider) Load() error {
	data, err := os.ReadFile(p.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return yaml.Unmarshal(data, p.config)
}

// Save 保存配置文件
func (p *Provider) Save() error {
	dir := filepath.Dir(p.configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(p.config)
	if err != nil {
		return err
	}

	return os.WriteFile(p.configPath, data, 0600)
}

// Get 获取配置
func (p *Provider) Get() *Config {
	return p.config
}

// Set 设置配置
func (p *Provider) Set(cfg *Config) {
	p.config = cfg
}

// SetServerURL 设置服务器地址
func (p *Provider) SetServerURL(url string) {
	p.config.ServerURL = url
}

// SetToken 设置 Token
func (p *Provider) SetToken(token string) {
	p.config.Token = token
}

// SetOutput 设置输出格式
func (p *Provider) SetOutput(output string) {
	p.config.Output = output
}

// ClearToken 清除 Token
func (p *Provider) ClearToken() {
	p.config.Token = ""
}

// GetConfigPath 获取配置文件路径
func (p *Provider) GetConfigPath() string {
	return p.configPath
}

// GetDefaultConfigPath 获取默认配置文件路径
func GetDefaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, DefaultConfigDir, DefaultConfigFile)
}

// GetConfigDir 获取配置目录
func GetConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, DefaultConfigDir)
}
