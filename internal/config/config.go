package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Listen      string   `yaml:"listen"`
	UpstreamDNS string   `yaml:"upstream_dns"`
	BlockList   []string `yaml:"block_list"`
	UseDoH      bool     `yaml:"use_doh"`
	DoHURL      string   `yaml:"doh_url"` // 例如 "https://dns.google/dns-query"
	DoHRequests []string `yamll:"doh_requests"`
	UseDoT      bool     `yaml:"use_dot"`
	DoTURL      string   `yaml:"dot_url"`
	DoTRequests []string `yaml:"dot_requests"`
	LogFile     string   `yaml:"log_file"`
	LogLevel    string   `yaml:"log_level"`
}

func LoadConfig(path string) *Config {
	cfg := &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
