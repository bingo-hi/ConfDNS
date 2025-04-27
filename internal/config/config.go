package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type LogConfig struct {
	FilePath   string `yaml:"filePath"`
	Level      string `yaml:"level"`
	MaxSizeMB  int    `yaml:"maxSizeMB"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAgeDays int    `yaml:"maxAgeDays"`
	Compress   bool   `yaml:"compress"`
}

type DNSConfig struct {
	Listen      string   `yaml:"listen"`
	UpstreamDNS string   `yaml:"upstream_dns"`
	BlockList   []string `yaml:"block_list"`
	UseDoH      bool     `yaml:"use_doh"`
	DoHURL      string   `yaml:"doh_url"`
	DoHRequests []string `yamll:"doh_requests"`
	UseDoT      bool     `yaml:"use_dot"`
	DoTURL      string   `yaml:"dot_url"`
	DoTRequests []string `yamll:"dot_requests"`
	CertFile    string   `yaml:"cert_file"`
	KeyFile     string   `yaml:"key_file"`
}

type AppConfig struct {
	LogConfig LogConfig `yaml:"log_config"`
	DNSConfig DNSConfig `yaml:"dns_config"`
}

func LoadAppConfigFromYAML(path string) (AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return AppConfig{}, err
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return AppConfig{}, err
	}
	return cfg, nil
}
