package main

import (
	"confdns/internal/config"
	"confdns/internal/dns"
	"confdns/internal/logger"
)

func main() {
	cfg := config.LoadConfig("config/config.yaml")
	logger.InitLogger(cfg.LogFile, cfg.LogLevel)
	dns.StartServer(cfg)
}
