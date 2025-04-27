package main

import (
	"confdns/internal/config"
	"confdns/internal/dns"
	"confdns/internal/logger"
	"fmt"
)

func main() {
	appCfg, err := config.LoadAppConfigFromYAML("config/config.yaml")
	if err != nil {
		fmt.Println("load config failed:", err)
	}

	logger.InitLoggerWithConfig(appCfg.LogConfig)
	dns.StartServer(appCfg.DNSConfig)
}
